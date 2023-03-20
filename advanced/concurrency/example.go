package main

import (
	"fmt"
	"time"
)

func main() {

	// 带close判断的

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for {
		if data, ok := <-ch; ok {
			time.Sleep(time.Second)
			fmt.Println(data)
		} else {
			break
		}
	}

	// 不带close判断的
	ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for i := 0; i <= 5; i++ {
		fmt.Println(<-ch)
		time.Sleep(time.Second)
	}

}
