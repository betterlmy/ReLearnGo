package main

import (
	"log"
	"sync"
	"time"
)

/*
一个将带缓冲channel用作计数信号量的例子
在这段代码中,jobs来表示总共需要完成的数量 activate来控制同时只能有多少个工作进行
所以说最终的结果是每三个工作是一个批次
*/

var activate = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {

	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			activate <- struct{}{} // 对空结构体传递一个空的字段
			log.Printf("处理工作:%d \n", j)
			time.Sleep(2 * time.Second)
			<-activate
			wg.Done()
		}(j)
	}

	wg.Wait()

}
