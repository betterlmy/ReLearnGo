go语言高级编程

## [接口](interface)

## 错误处理

### [error](error)

Go语言中没有异常机制,但是有一个error接口,可以通过error接口来处理错误.

#### 自定义error
除了使用errors包中的New方法来创建自定义异常,还可以通过接口来创建.  
Go语言中的error类型实际上是一个接口,只要实现了Error() string方法即可,源代码如下:
```go
type error interface{
	Error() string
}
```
自定义一个类型,实现返回string类型的Error()方法即可,代码如下:
```go
type ErrNegSqrt float64
func (e ErrNegSqrt) Error() string{
	return fmt.Sprintf("不能对负数`%v`进行开根号",e)
}
```
[完整调用代码](error/error/error.go)

### [panic宕机](error/panic)

一般而言,只有当程序发生不可逆的错误时,才需要使用`panic`方法来触发宕机.  
panic方法是Go语言的内置函数,使用该方法后,程序将直接中断.
panic的源代码如下,可以看到panic方法接收一个interface{}类型的参数,该参数可以是任意类型.

```go
func panic(v interface{})
```

#### 调用panic的情形:

* 程序处于失控状态且无法恢复,继续执行会影响其他程序时
* 发生不可预知的错误时
* 程序运行到不应该运行到的地方时

### [recover宕机恢复](error/recover)

Go语言通过内置函数recover来捕获宕机,类似于其他语言的try-catch机制  
在使用panic方法触发宕机之后,且在退出当前函数前,会延迟调用defer语句

```go
package main

import "fmt"

func protect() {
	defer func() {
		fmt.Println("func protect 退出")
	}()
	panic("panic in protect")
}

func main() {
	defer func() {
		fmt.Println("func main 退出")
	}()
	protect()
	fmt.Printf("这段代码无法被执行")
}

//输出
//func protect 退出
//func main 退出
//panic: panic in protect

```

由于defer语句低延迟执行的特性,我们可以通过`defer语句+匿名函数+recover方法`来实现对宕机的捕获,代码如下:

```go
package main

import "fmt"

func protect() {
	defer func() {
		if err := recover(); err != nil { // 可以获取到panic传入的参数
			fmt.Println("recover from ", err)
		}
	}()
	panic("panic in protect")
}

func main() {
	protect()
	fmt.Printf("这段代码无法被执行")
	// 执行结果
	//recover from  panic in protect
	//这段代码无法被执行

}
```

上述代码中,在protect函数中,首先defer了一个匿名函数,这个函数主要用于panic抓取,当panic产生了宕机之后,并不会中断整个程序的运行,而是进入到了defer语句中,而defer语句中的匿名函数做到了panic的恢复工作recover,且recover可以获取panic传入的参数,所以说整个main()
函数足以全部执行

实际应用中,我们可以通过写入安全函数的方式,即使发生panic 主程序仍然能够继续正常运行,[示例代码](error/recover/recover.go)

