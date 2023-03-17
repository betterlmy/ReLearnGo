package main

import (
	"fmt"
	"sync"
	"time"
)

// produce 生产者 只发送数据
func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

// consume 消费者 只接受数据
func consume(ch <-chan int) {
	for n := range ch {
		fmt.Println("接收到的数据", n)
	}
}

// sendOnlyAndReceiveOnly 只发送和只接受
func sendOnlyAndReceiveOnly() {

	ch3 := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch3)
		fmt.Println("wg done1")
		wg.Done()
	}()

	go func() {
		consume(ch3)
		fmt.Println("wg done2")
		wg.Done()
	}()

	fmt.Println("wg waiting")
	wg.Wait()
	fmt.Println("程序运行结束")

}

func InitAndSendAndReceive() {
	// 如果channel类型变量在声明的时候没有被赋予初值,那么他默认值为nil
	// 并且和其他复合数据类型支持使用复合类型字面值做为变量初始值不一样,channel 只能使用make方法进行预定义
	ch1 := make(chan int)    // 无缓冲channel
	ch2 := make(chan int, 2) // 带缓冲长度为2的channel类型 这些channel只能够存放int数据

	// 发送与接受
	go func() {
		// 无缓冲区的channel只能够在接受者准备好的时候才能够发送数据
		ch1 <- 2
	}()
	// 带缓冲 channel 的发送操作在缓冲区未满、接收操作在缓冲区非空的情况下是异步的（发送或接收不需要阻塞等待）。
	ch2 <- 17
	ch2 <- 17
	//ch2 <- 17 // fatal error: all goroutines are asleep - deadlock! 因为长度为2 此时已经传入了过多的数据

	//go func() {
	//	ch2 <- 5
	//}()
	m := <-ch2

	n := <-ch2
	fmt.Println(m, n)
	fmt.Println(<-ch1)
}

type signal struct{}

func worker() {
	time.Sleep(1 * time.Second)
	println("worker is working...")
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}
