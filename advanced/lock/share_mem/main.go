package main

import (
	"fmt"
	"sync"
)

func main() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			mut.Lock()
			defer mut.Unlock()
			count++
			wg.Done()
		}()
	}
	wg.Wait() // 等待所有wg中的协程都执行完毕 会阻塞
	fmt.Println(count)
}
