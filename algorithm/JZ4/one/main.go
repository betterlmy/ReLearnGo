package main

import "fmt"

/*
 * 13ms4932kb
 */

func Find(target int, array [][]int) bool {
	// write code here

	maxWidth := len(array) - 1
	if maxWidth < 0 {
		return false
	}
	fmt.Println("maxWidth", maxWidth)
	for i1, ints := range array {

		for i2, value := range ints {
			if i2 > maxWidth {
				fmt.Println(i1, "行", i2, "列pass")
				break
			}
			if value == target {
				return true
			} else if value > target {
				maxWidth = i2 - 1
				fmt.Println("maxWidth", maxWidth)
				if maxWidth < 0 {
					return false
				}
				break
			}
		}
	}
	return false
}

func test(array [][]int, target int) {
	for _, ints := range array {
		fmt.Println(ints)
	}
	fmt.Println(target)
	fmt.Print(Find(target, array))
}

func main() {
	array1 := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	target1 := 7

	target2 := 3

	test(array1, target1)
	test(array1, target2)
}
