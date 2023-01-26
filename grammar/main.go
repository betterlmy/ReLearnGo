package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
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
func eval(a, b int, op string) (float64, error) {
	var result float64
	var err error = nil
	switch op {
	case "+":
		result = float64(a + b)
	case "-":
		result = float64(a - b)
	case "*":
		result = float64(a * b)
	case "/":
		if b == 0 {
			err = fmt.Errorf("除数不能为0")
		}
		result = float64(a) / float64(b)
	default:
		err = fmt.Errorf("不支持的操作子:%s", op)
	}
	return result, err
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

// increase 学习for循环
func increase() int {
	var sum int
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	return sum
}

// convert2Bin 十进制转二进制,学习for循环
func convert2Bin(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for ; n > 0; n /= 2 {
		x := n % 2
		// 从 int 到 string 的转换将整数值解释为代码点
		result = strconv.Itoa(x) + result
	}
	return result
}

// forever 无限循环,while True语句
func forever() {
	for {
		fmt.Println("abc")
	}
}

// div 带余数的除法,学习函数中的双返回值,给予命名可以在调用时候返回该命名,方便使用
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

// apply 将函数作为参数使用
func apply(op func(float64, float64) float64, a, b float64) float64 {
	p := reflect.ValueOf(op).Pointer()    //返回指向函数的指针
	opName := runtime.FuncForPC(p).Name() //返回函数名
	fmt.Println("执行函数:", opName, "with args", a, " ", b)
	return op(a, b)
}

// sum 求和,学习可变参数
func sum(num ...int) int {
	s := 0
	for _, j := range num {
		s += j
	}
	return s
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
	//fmt.Println(grade(0))
	//fmt.Println(increase())
	//fmt.Println(convert2Bin(12344))
	//io.PrintFile("abc.txt")
	//forever()
	//fmt.Println(div(5, 3))

	//f, err := eval(5, 1, "/")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(f)
	//fmt.Println(apply(math.Pow, 3, 4))
	//fmt.Println(sum(1, 2, 3))
	//pointer()
	pointerTypeCheck()
	//arr1, arr2, arr3 := arrays()
	//printArray(arr1)
	//printArray(arr2) // 会报错,因为数组是值类型,传递的是值的拷贝
	//printArray(arr3)
}
