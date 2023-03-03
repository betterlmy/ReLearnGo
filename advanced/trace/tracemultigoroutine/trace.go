package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

/*
当Trace函数面对一个Goroutine时还是可以支撑的,但是当程序中并发运行多个goroutine时,多个函数调用链的出入口信息就会混杂在一起.
解决方案:
	在输出的函数出入口信息时,单上一个在程序每次执行时能够唯一区分goroutine的ID
*/

var goroutineSpace = []byte("goroutine ")

// curGoroutineID 获取GoroutineID
func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

// Trace 根据goroutine获取函数信息
func Trace() func() {
	pc, _, _, ok := runtime.Caller(1) //获取当前goroutine上的调用栈信息 skip=0返回调用者的函数信息,skip返回调用者的信息
	// 四个返回值,pc程序技术,原文件名,所在行数,返回是否成功
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)

	gid := curGoroutineID()
	name := fn.Name()
	fmt.Printf("g[%05d]: Enter: [%s]\n", gid, name)
	// ID不足五位数则左补0
	return func() {
		fmt.Printf("g[%05d]: Exit: [%s]\n", gid, name)
	}

}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()

}

func A1() {
	defer Trace()()
	B1()
}

func B1() {
	defer Trace()()
	C1()
}

func C1() {
	defer Trace()()
	D()
}

func D() {
	defer Trace()()
}

func A2() {
	defer Trace()()
	B2()
}

func B2() {
	defer Trace()()
	C2()
}

func C2() {
	defer Trace()()
	D()
}
