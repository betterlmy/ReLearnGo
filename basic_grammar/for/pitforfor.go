package main

import (
	"fmt"
	"time"
)

func main() {
	var m = []int{1, 2, 3, 4, 5}

	for i, v := range m {
		fmt.Println(i, v)
		go func(int, int) {
			fmt.Println(i, v)
		}(i, v)

		go func() {
			fmt.Println(i, v)
			// 会出现异常,因为i,v在循环中地址并不会发生变化,每次的循环都会导致i,v的内容发生变化,
			// 在等待了三秒钟之后,每个goroutine都会获取同一个地址的值,所以说输出是相同的
		}()

	}

	time.Sleep(time.Second * 10)
}
