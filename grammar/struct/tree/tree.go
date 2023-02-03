package tree

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Left, Right *Tree
	Value       int
}

// Print 限定接受体的方法
func (node *Tree) Print() {
	fmt.Print(strconv.Itoa(node.Value) + " ")
	fmt.Println()
}
func (node *Tree) SetValue(value int) {
	if node == nil {
		fmt.Println("node不存在")
	} else {
		node.Value = value
	}
}

func (node *Tree) TraverseMid() {
	// 中序遍历
	if node == nil {
		return
	}
	node.Left.TraverseMid()
	node.Print()
	node.Right.TraverseMid()
}

func Init() {
	var root1 Tree
	fmt.Println(root1)

	root2 := Tree{Left: &root1, Value: 1}
	root2.Right = &Tree{nil, nil, 5}
	fmt.Println(root2)
	fmt.Println(root2.Right.Value)
	root2.Print()
	root3 := Tree{}
	root3.SetValue(3)
	fmt.Println(root3)
	fmt.Println("遍历")
	root2.Right.Left = &root3
	root2.TraverseMid()

}
