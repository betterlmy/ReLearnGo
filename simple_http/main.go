// The code creates a channel, ch, and two goroutines.
// The first goroutine creates a channel, ch, and sends a 1 to it.
// The second goroutine receives from the channel, ch, after a second.
// The main goroutine prints the number of goroutines in the program every second.

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		// print the number of goroutines in the program
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
