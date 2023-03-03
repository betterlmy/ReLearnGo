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
	p1 := Teacher{
		Person: Person{Title: "教授"},
		Name:   "李",
		Title:  "先生",
	}
	println(p1.Name)
	println(p1.Person.Title)
}

func hello() {
	fmt.Println("Hello World")
}
