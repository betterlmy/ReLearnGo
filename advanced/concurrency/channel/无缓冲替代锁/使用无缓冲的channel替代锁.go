package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	c chan int
	i int
}

func NewCounter() *Counter {
	cter := &Counter{
		c: make(chan int), // 使用无缓冲channel
		i: 0,
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *Counter) Increase() int {
	// 因为无缓冲的channel只有在接受之后 下一个goroutine才能够发送
	return <-cter.c
}

func main() {
	cter := NewCounter() // 创建一个线程 无线增加计数器 但是只有对channel接受一次 for循环才能执行一次
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}

	wg.Wait()

}
