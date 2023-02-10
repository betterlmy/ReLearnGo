package main

import "fmt"

func pointer() {
	var a int = 2
	var pointerA *int = &a
	*pointerA = 3
	fmt.Println("指向A的内容为", a)
}

func passByValue(a int) {
	a += 1
}

func passByPointer(a *int) {
	*a += 1
}

func pointerTypeCheck() {
	a, b := 3, 3
	passByValue(a)
	passByPointer(&b)
	fmt.Println("值传递", a)
	fmt.Println("指针传递", b)
}
