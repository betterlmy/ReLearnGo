package main

import "fmt"

/*
输入一个链表的头节点，按链表**从尾到头**的顺序返回每个节点的值（用数组返回）
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

var Out = make([]int, 10000)
var Length = 0

func printListFromTailToHead(head *ListNode) []int {
	/*
		递归
	*/
	if head == nil {
		return []int{}
	}

	getNodeValue(head)
	Out[Length] = head.Val

	return Out[:Length+1]

}

func getNodeValue(ln *ListNode) int {
	if ln.Next != nil {
		Out[Length] = getNodeValue(ln.Next) // 递归 最后一个没有next的节点的length=0
		Length++
	}
	return ln.Val
}

func main() {
	var ln3 = ListNode{
		Val:  3,
		Next: nil,
	}

	var ln2 = ListNode{
		Val:  2,
		Next: &ln3,
	}

	var ln1 = ListNode{
		Val:  1,
		Next: &ln2,
	}

	fmt.Print(printListFromTailToHead(&ln1))
}
