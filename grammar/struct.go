package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Left, Right *TreeNode
	Value       int
}

// 限定接受体的方法
func (node *TreeNode) print() {
	fmt.Print(strconv.Itoa(node.Value) + " ")
}
func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("node不存在")
	} else {
		node.Value = value
	}
}

func (node *TreeNode) traverseMid() {
	// 中序遍历
	if node == nil {
		return
	}
	node.Left.traverseMid()
	node.print()
	node.Right.traverseMid()
}
func structInit() {
	var root1 TreeNode
	fmt.Println(root1)

	root2 := TreeNode{Left: &root1, Value: 1}
	root2.Right = &TreeNode{nil, nil, 5}
	fmt.Println(root2)
	fmt.Println(root2.Right.Value)
	root2.print()
	root3 := TreeNode{}
	root3.setValue(3)
	fmt.Println(root3)
	fmt.Println("遍历")
	root2.Right.Left = &root3
	root2.traverseMid()

}
