package main

import "fmt"

/*
输入一个链表的头节点，按链表**从尾到头**的顺序返回每个节点的值（用数组返回）
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	/*
		先获取正向,再将[]int反向打印
	*/
	var out []int
	if head == nil {
		return out
	}
	for {
		out = append(out, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next

	}
	return reverse(out)

}

func reverse(input []int) []int {
	lens := len(input)
	out := make([]int, lens)
	for i := lens; i > 0; i-- {
		println(lens-i, i)
		out[lens-i] = input[i-1]
	}
	return out
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
