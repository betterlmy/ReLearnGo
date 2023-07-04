package main

import (
	"fmt"
)

type Person struct {
	Title string
}
type Teacher struct {
	Person
	Name  string
	Title string
}

func main() {
	//hello()
	// p1 := Teacher{
	// 	Person: Person{Title: "教授"},
	// 	Name:   "李",
	// 	Title:  "先生",
	// }
	// println(p1.Name)
	// println(p1.Person.Title)

	// x := "你好,世界"
	x := []rune("你好")
	x = []rune("good你好")
	_ = x
}

func hello() {
	fmt.Println("Hello World")
}
