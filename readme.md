# 前言

编写时间:2023-01-11

根据[李文周老师课程](https://www.liwenzhou.com/)进行学习.  
Golang视频,笔记,学习路线等资料[链接](https://www.aliyundrive.com/s/yUJBxdZxFk8) 提取码:88dk  
在go1.5之前需要使用GOPATH进行工作环境配置,最新已经不需要使用了,新的HelloWorld视频可以看[链接](https://www.bilibili.com/video/BV1bV41177KD)

go1.18版本添加了go work工作区模式

## 学习顺序

### [输出Hello World](./helloworld)

### [简单的gin_demo](./gin_demo)

### [go_module依赖管理工具](./go_module)

### [基础语法](./basic_grammar)

### [高级编程](./advanced)

## Go语言的卖点

* 作为一门静态语言,却和很多动态脚本语言一样灵活
* 节省内存,程序启动快,执行速度快,编译速度快
* 内置并发编程
* 良好的代码可读性
* 跨平台支持
* 活跃的社区

## Go语言的特性

* 内置并发编程
    * 使用协程(goroutine)作为基本的计算单元.轻松创建协程
    * 使用通道(channel)来实现协程见的同步和通信
* 内置[映射(map)](basic_grammar/map.go)和[切片(slice)](basic_grammar/slice.go)类型
* 支持多态(polymorphism)
* 使用[接口(interface)](./advanced/interface)来实现装盒(value boxing)和反射(reflection)
* 支持[指针](basic_grammar/pointer.go)
* 支持函数闭包(closure) 
* 支持方法
* 支持延迟函数调用(defer)
* 支持[类型内嵌](./basic_grammar/tree2.go)
* 支持类型推断(type deduction or type inference)
* 类型安全
* 自动垃圾回收
* 跨平台
* 自定义泛型(从Go 1.18开始)


## 一些链接

* [中文标准库1](https://studygolang.com/pkgdoc)
* [中文标准库2](https://github.com/astaxie/gopkg)
* [Golang官网1](https://go.dev/)
* [Golang官网2](https://golang.google.cn/)  
* [Github资源](https://github.com/avelino/awesome-go)


