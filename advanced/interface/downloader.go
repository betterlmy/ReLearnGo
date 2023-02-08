package main

import (
	"fmt"
	"interface/infra"
	"interface/testing"
)

type retriever interface {
	//接口是个抽象的概念,只告诉这个接口包含了什么方法,不包含方法的实现

	Get(string) error // 这个接口包含着Get这个方法,同时限定了方法的参数和返回值
}

func main() {
	url := "https://www.google.com/robots.txt"

	var retriever1 retriever = infra.Retriever{}
	var retriever2 retriever = testing.Retriever{}

	err1 := retriever1.Get(url)
	err2 := retriever2.Get(url)
	if err1 == nil && err2 == nil {
		fmt.Println("下载完成")
	}
}
