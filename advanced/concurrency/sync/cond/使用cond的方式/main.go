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

func spawnGroup(f func(i int), num int, groupSignal *sync.Cond) <-chan signal {
	// 生成一组工人开始工作 不过开始工作前需要不断的轮询查询是否可以开始工作
	c := make(chan signal) // 无缓冲的channel 同步
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(workerIndex int) {
			groupSignal.L.Lock()
			for !ready {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			f(workerIndex)
			wg.Done()
			return
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

	groupSignal := sync.NewCond(&sync.Mutex{}) // 为添加一个条件信号

	c := spawnGroup(worker, 5, groupSignal)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of worker start to work")

	groupSignal.L.Lock()
	ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("all workers work done")

}
