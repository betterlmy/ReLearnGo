package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func spawn(f func(*chan error, int)) []chan error {
	chans := make([]chan error, 3)
	for i := 0; i < 3; i++ {
		chans[i] = make(chan error)
		go f(&chans[i], i)
	}
	return chans
}

func setErr(c *chan error, i int) {
	time.Sleep(2 * time.Second)
	// 必须对所有的channel都进行写入操作
	if i == 0 {
		*c <- nil
	}
	err := errors.New("第" + strconv.Itoa(i) + "个goroutine发生了一个错误")
	*c <- err

}

func main() {
	chans := spawn(setErr)
	time.Sleep(time.Second * 3)
	for _, c := range chans {
		if c == nil {
			println("c")
		}
		fmt.Println(<-c)
	}

}
