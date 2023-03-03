package main

import "runtime"

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1) //获取当前goroutine上的调用栈信息 skip=0返回调用者的函数信息,skip返回调用者的信息
	// 四个返回值,pc程序技术,原文件名,所在行数,返回是否成功
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	println("Enter:", name)
	return func() { println("Exit:", name) }

}

func main() {
	defer Trace()()
	foo()
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}
