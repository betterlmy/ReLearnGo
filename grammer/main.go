package main

import "fmt"

func main() {
	fmt.Println("Hello Grammer!")

	var str string = "helloworld"
	var age int = 5
	des := "这是一个描述"
	fmt.Print(str + "\n")
	fmt.Printf("我要输出%s,age=%d,描述=%s\n", str, age, des)
}
