package main

import (
	"fmt"
	"sync"
)

/*
无缓冲channel具有同步特性,让他在某些场合下可以替代锁
在这个示例中，我们使用了一个带有互斥锁保护的全局变量作为计数器，所有要操作计数器的 Goroutine 共享这个全局变量，并在互斥锁的同步下对计数器进行自增操作。
*/

type Counter struct {
	sync.Mutex
	i int
}

var counter Counter

func Increase() int {
	counter.Lock()
	defer counter.Unlock()
	counter.i++
	return counter.i
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine:%d:current counter value is :%d \n", i, v)
			wg.Done()
		}(i)
	}

	wg.Wait()

}
