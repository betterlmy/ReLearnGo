package main

import (
	"fmt"
	"math"
	"reflect"
)

// variable 学习变量类型
func variable() {
	// 变量的声明和定义
	var str1 string = "hello"        //标准定义
	var str2 = "-world2"             //隐式定义
	age, time, des := 5, 7, "这是一个描述" //短变量声明,加类型推断
	fmt.Print(str1+str2, "\n")
	fmt.Printf("我要输出%s,age=%d,time=%d,描述=%s\n", str1, age, time, des)
}

// triangle 学习变量类型转换
func triangle() {
	a, b := 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	d := int(c)
	fmt.Printf("c的类型 %T\n", c)
	fmt.Println("d的类型", reflect.TypeOf(d))
}

// consts 学习常量类型
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
	fmt.Println("a的类型", reflect.TypeOf(a))
}

// enums 学习枚举类型
func enums() {
	const (
		cpp = iota
		java
		python
	)
	fmt.Println(cpp, java, python)

	//b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Printf("1kb=%db\n", kb)
	fmt.Printf("1mb=%db\n", mb)
	fmt.Printf("1gb=%db\n", gb)
	fmt.Printf("1tb=%db\n", tb)
	fmt.Printf("1pb=%db\n", pb)

}

// bounded 判断一个数是否在0-100之间学习if语句
func bounded(v int) int {
	if v >= 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

// eval 用于计算加减乘除,学习switch语句
func eval(a, b int, op string) float64 {
	var result float64
	switch op {
	case "+":
		result = float64(a + b)
		fallthrough
	case "-":
		result = float64(a - b)
	case "*":
		result = float64(a * b)
	case "/":
		if b == 0 {
			panic("除数不能为0")
		}
		result = float64(a) / float64(b)
	default:
		panic("不支持的操作子:" + op)
	}
	return result
}

// grade 用于分数评估,学习switch中添加比较语句
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return "得分为" + g
}
func main() {
	fmt.Println("Hello Go Grammar!")
	//variable()
	//triangle()
	//consts()
	//enums()
	//for _, j := range []int{-10, 20, 100, 100} {
	//	fmt.Println(bounded(j))
	//}
	//io.ReadFileContent("abc.txt")
	//fmt.Println(eval(1, 1, "+"))
	fmt.Println(grade(0))
}
