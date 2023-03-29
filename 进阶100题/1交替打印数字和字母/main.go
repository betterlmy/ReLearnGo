package main

import (
	"fmt"
	"sync"
	"time"
)

/*
使用Goroutine 交替打印数组和字母
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

思路 设置两个无缓冲的channel，一个用来打印数字，一个用来打印字母
*/

func printLetter(chNum, chLetter chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ascii := 65

	for {
		timer := time.NewTimer(time.Millisecond)
		select {
		case <-chLetter:
			fmt.Printf("%c%c", ascii, ascii+1)
			ascii += 2
			chNum <- struct{}{}
		case <-timer.C:
			return
		}
	}
}

func printNum(chNum, chLetter chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		select {
		case <-chNum:
			fmt.Printf("%d%d", i, i+1)
			i += 2
			if i >= 29 {
				return
			} else {
				chLetter <- struct{}{}
			}
		}
	}

}

func main() {

	var chNum, chLetter = make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go printNum(chNum, chLetter, &wg)
	go printLetter(chNum, chLetter, &wg)
	chNum <- struct{}{}

	wg.Wait()
	close(chNum)
	close(chLetter)

}
