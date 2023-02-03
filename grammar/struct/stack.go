package Struct

import "fmt"

type Stack struct {
	Data []int
}

//type Stack []int

func (q *Stack) Push(v int) {
	(*q).Data = append((*q).Data, v)
}

func (q *Stack) Pop() int {
	tail := (*q).Data[len((*q).Data)-1]
	(*q).Data = (*q).Data[:len((*q).Data)-1]
	return tail
}

func (q *Stack) IsEmpty() bool {
	return len((*q).Data) == 0
}

func (q *Stack) Print() {
	if len((*q).Data) == 0 {
		return
	}
	for _, v := range (*q).Data {
		fmt.Print(v)
	}
	fmt.Println()

}
