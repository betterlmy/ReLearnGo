package stack

//

import (
	"errors"
	"fmt"
)

//type IntStack interface {
//	Push(int) error
//	Pop() (int error)
//}

type intStack struct {
	data []int64
	Len  int
	Cap  int
}

func (stack *intStack) Push(input int64) error {
	if stack.Len == stack.Cap {
		// 如果支持动态栈 删掉这个代码即可
		return errors.New("栈满")
	}

	stack.data = append(stack.data, input)

	if stack.Len == stack.Cap {
		stack.Cap = cap(stack.data) // 如果支持动态栈
		fmt.Println("栈已扩容")
	}
	stack.Len += 1
	fmt.Printf("push元素%d成功,当前栈的容量%d,当前长度%d\n", input, stack.Cap, stack.Len)
	return nil
}
func (stack *intStack) Pop() (int64, error) {
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
		data: make([]int64, 0, cap),
		Len:  0,
		Cap:  cap,
	}, nil
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
