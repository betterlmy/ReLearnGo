package Struct

import "fmt"

type Queue struct {
	Data []int
}

//type Stack []int

func (q *Queue) Push(v int) {
	(*q).Data = append((*q).Data, v)
}

func (q *Queue) Pop() int {
	head := (*q).Data[0]
	(*q).Data = (*q).Data[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len((*q).Data) == 0
}

func (q *Queue) Print() {
	if len((*q).Data) == 0 {
		return
	}
	for _, v := range (*q).Data {
		fmt.Print(v)
	}
	fmt.Println()
}
