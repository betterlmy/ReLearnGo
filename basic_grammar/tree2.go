package main

// tree2文件中,MyTree通过内嵌的方式进行
import (
	"fmt"
	"grammar/struct/tree"
	"strconv"
)

type MyTree struct {
	// 通过内嵌的方式进行继承
	*tree.Tree
}

func (node MyTree) Print() {
	if node.Tree != nil {
		fmt.Print(strconv.Itoa(node.Value) + " ")
	}
}

func (node MyTree) SetValue(v int) {
	if node.Tree == nil {
		fmt.Println("node不存在")
	} else {
		node.Value = v
	}
}

func (node MyTree) TraversePre() {
	if node.Tree == nil {
		return
	}
	fmt.Print(strconv.Itoa(node.Value) + " ")
	left := MyTree{node.Left}
	left.TraversePre()
	right := MyTree{node.Right}
	right.TraversePre()
}
