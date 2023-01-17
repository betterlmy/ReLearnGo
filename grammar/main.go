package main

import (
	"fmt"
	"math"
	"reflect"
)

func variable() {
	// 变量的声明和定义
	var str1 string = "hello"        //标准定义
	var str2 = "-world2"             //隐式定义
	age, time, des := 5, 7, "这是一个描述" //短变量声明,加类型推断
	fmt.Print(str1+str2, "\n")
	fmt.Printf("我要输出%s,age=%d,time=%d,描述=%s\n", str1, age, time, des)
}

func triangle() {
	a, b := 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	d := int(c)
	fmt.Printf("c的类型 %T\n", c)
	fmt.Println("d的类型", reflect.TypeOf(d))
}

func main() {
	fmt.Println("Hello Grammar!")
	//variable()
	triangle()
}
