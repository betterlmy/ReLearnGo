package main

import (
	"datastruct/stack"
	"fmt"
)

func main() {
	intStack, _ := stack.NewIntStack(5)
	_ = intStack.Push(5)

	pop, err := intStack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println(pop)

}
