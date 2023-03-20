package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct {
}

var ready = false

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d: is work done...\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	// 生成一组工人开始工作 不过开始工作前需要不断的轮询查询是否可以开始工作
	c := make(chan signal) // 无缓冲的channel 同步
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(workerIndex int) {
			for {
				// 创建一个for循环 不断的轮询是否可以开始工作
				mu.Lock() // 尝试加锁
				if !ready {
					mu.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mu.Unlock()
				f(workerIndex)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		// 打开一个协程用来判断是否全部已经工作完成
		wg.Wait()
		c <- signal{}
	}()
	return c

}

func main() {
	fmt.Println(" start a group of workers...")

	mu := &sync.Mutex{}
	c := spawnGroup(worker, 5, mu)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of worker start to work")

	mu.Lock()
	ready = true
	mu.Unlock()

	<-c
	fmt.Println("all workers work done")

}
