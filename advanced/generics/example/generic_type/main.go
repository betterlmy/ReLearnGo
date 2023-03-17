package main

import "fmt"

type Vector[T any] []T

func (v Vector[T]) Dump() {
	fmt.Printf("%v\n", v)
}

func main() {
	x := Vector[string]{"你好", "Hello"}
	y := Vector[int]{1, 2, 3}

	x.Dump()
	y.Dump()

}
