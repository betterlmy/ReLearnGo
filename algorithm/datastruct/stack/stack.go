package stack

//

import (
	"errors"
	"fmt"
)

type Stack interface {
	push(int) error
	pop() (int error)
}

type intStack struct {
	data []int
	Len  int
	Cap  int
}

func (stack *intStack) Push(input int) error {
	if stack.Len == stack.Cap {
		return errors.New("栈满")
	}
	stack.data = append(stack.data, input)
	stack.Len += 1
	fmt.Printf("push元素%d成功,当前栈的容量%d,当前长度%d\n", input, stack.Cap, stack.Len)
	return nil
}
func (stack *intStack) Pop() (int, error) {
	if stack.Len == 0 {
		return -1, errors.New("栈空")
	}
	tmp := stack.data[stack.Len-1]
	stack.data = stack.data[:stack.Len-1]
	stack.Len -= 1
	fmt.Printf("pop出%d成功,当前栈的容量%d,当前长度%d\n", tmp, stack.Cap, stack.Len)
	return tmp, nil
}

// NewIntStack 设置一个方法返回一个栈
func NewIntStack(cap int) (*intStack, error) {
	if cap <= 0 {
		return nil, errors.New("cap应该大于0")
	}

	fmt.Printf("新建stack成功,cap=%d\n", cap)

	return &intStack{
		data: make([]int, 0, cap),
		Len:  0,
		Cap:  cap,
	}, nil
}
func Test() {
	fmt.Println("1")
}

func main() {
	stack, err := NewIntStack(1)
	if err != nil {
		panic(err)
	}

	err = stack.Push(2)
	if err != nil {
		panic(err)
	}

	_, err = stack.Pop()
	if err != nil {
		panic(err)
	}

}
