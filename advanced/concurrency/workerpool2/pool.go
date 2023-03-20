package main

/*
添加功能选项机制.
*/

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type Option func(*Pool)
type Task func()

type Pool struct {
	cap int //最大容量

	active chan struct{}  // 计数信号量 有缓冲,表示当前active的worker线程数量,每个worker尝试向 active 中传入一个空的结构体,如果阻塞则表明没有空闲,则继续等待
	tasks  chan Task      // task channel 无缓冲 用来存储任务,每个worker会从 tasks 这个通道中来获取自己将来要做的任务
	wg     sync.WaitGroup // 用于在 Pool 销毁时等待所有worker退出
	quit   chan struct{}  // 通知各个worker退出的信号channel 无缓冲当需要停止 Goroutine 池中的所有 worker 并退出时，向 quit channel 发送一个空的结构体，所有 worker 线程都会在处理完当前任务后立即退出。

	preAlloc bool // 是否在创建pool的时候就预创建workers，默认值为：false
	block    bool // 当pool满的情况下，新的Schedule调用是否阻塞当前goroutine。默认值：true,如果block = false，则Schedule返回ErrNoWorkerAvailInPool
}

// WithBlock 添加一个功能选项，用于设置当pool满的情况下，新的Schedule调用是否阻塞当前goroutine。
func WithBlock(block bool) Option {
	return func(p *Pool) {
		p.block = block
	}
}

// WithPreAllocWorkers 添加一个功能选项，用于设置是否在创建pool的时候就预创建workers
func WithPreAllocWorkers(preAlloc bool) Option {
	return func(p *Pool) {
		p.preAlloc = preAlloc
	}
}

var ErrWorkerPoolFreed = errors.New("workerpool freed") // workerpool已终止运行

func (p *Pool) run() {
	idx := len(p.active)

	if !p.preAlloc {
		// 如果没有预创建workers，则根据任务来创建workers
	loop:
		for t := range p.tasks {
			p.returnTask(t)
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				p.newWorker(idx)
			default:
				break loop
			}
		}
	}

	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			// 创建新的worker
			idx++
			p.newWorker(idx)
		}
	}
}

// newWorker 生成一个worker开始工作 并进行判断是否停止工作
func (p *Pool) newWorker(workerId int) {
	p.wg.Add(1)
	go func() {
		// 为工人打开一个线程，用来监听下来是否有任务 这个线程中单独保存着
		defer func() {
			// 防止提交的task抛出panic
			if err := recover(); err != nil {
				log.Printf("[worker %d]: panic:%s\n", workerId, err)
				<-p.active
			}
			p.wg.Done()

		}()

		log.Printf("[worker %d]: start\n", workerId)

		for {
			// 线程池的关键, 通过不断地select来判断是否有任务,而不是释放线程
			select {
			case <-p.quit:
				log.Printf("[worker %d]: exit\n", workerId)
				<-p.active
				return
			case t := <-p.tasks:
				log.Printf("[worker %d]: recived a task\n", workerId)
				t()
				log.Printf("[worker %d]: finished a task\n", workerId)

			}
			time.Sleep(10 * time.Millisecond)
		}

	}()
}

// Schedule 提供给用户的提交请求的导出方法 每次提交一个任务给worker
func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil

	default:
		if p.block {
			p.tasks <- t
			return nil
		}
		return ErrWorkerPoolFreed
	}

}

// Free 释放线程池
func (p *Pool) Free() {
	close(p.quit) // 对于无缓冲的channel quit,close方法执行之后 <-p.quit会立即返回并不会堵塞
	p.wg.Wait()
	log.Println("[worker pool]:freed")
}

func (p *Pool) returnTask(t Task) {

}

var defaultCap = 5
var maxCap = 10

// New 实例化线程池
func New(cap int, opts ...Option) *Pool {

	if cap <= 0 {
		// 传入值校验
		cap = defaultCap
	} else if cap > maxCap {
		cap = maxCap
	}

	p := &Pool{
		cap:    cap,
		active: make(chan struct{}, cap),
		tasks:  make(chan Task),
		wg:     sync.WaitGroup{},
		quit:   make(chan struct{}),
	}

	for _, opt := range opts {
		// 通过遍历opts来设置Pool的属性
		opt(p)
	}
	log.Printf("workerpool start(preAlloc=%t)\n", p.preAlloc)
	if p.preAlloc {
		// 创建所有的worker
		for i := 0; i < p.cap; i++ {
			p.newWorker(i + 1)
			p.active <- struct{}{}
		}

	}
	go p.run()
	return p
}

func main() {
	pool := New(5, WithBlock(true), WithPreAllocWorkers(true))
	time.Sleep(time.Second * 1) // 等待线程池创建完成(此时)
	for i := 0; i < 10; i++ {
		fmt.Println("这是一个新的任务:", i)
		err := pool.Schedule(func() {
			// 将这个匿名任务传入到pool方法提供的接口中
			time.Sleep(time.Second * 1)
		})
		if err != nil {
			println("task: ", i, "err:", err)
		}
	}

	pool.Free()

}
