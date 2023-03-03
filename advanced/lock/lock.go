package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func increment() {
	// 区别于C语言,mutex对象并不需要进行init操作
	mutex.Lock()
	defer mutex.Unlock()
	count++
}

func main() {
	var wg sync.WaitGroup // 什么含义?
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}
