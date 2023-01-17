package io

import (
	"fmt"
	"os"
	"reflect"
)

func ReadFileContent(filename string) {
	// ioutil.ReadFile已在1.16版本被去除,使用os包进行ReadFile操作
	dir, _ := os.Getwd()
	fmt.Println("当前工作目录:", dir)

	// 条件语句写法1
	contents1, err1 := os.ReadFile(filename)
	if err1 != nil {
		fmt.Println("读取", filename, "失败,", err1)
	} else {
		fmt.Println(reflect.TypeOf(contents1), "\n", string(contents1))
		// contents1为[]uint8类型,所以需要转换为string类型
	}
	fmt.Println("程序继续执行")

	// 条件语句写法2,条件内赋值,赋值作用域仅限于if语句内
	if contents2, err2 := os.ReadFile(filename); err2 != nil {
		fmt.Println("读取", filename, "失败,程序中断运行")
		panic(err2) // panic会中断多线程的执行
	} else {
		fmt.Println(reflect.TypeOf(contents2), "\n", contents2)
	}
	fmt.Println("程序继续执行")

}
