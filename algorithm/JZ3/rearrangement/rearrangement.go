package main

import "fmt"

// duplicate 给定一个数组,如果出现了重复的数字,则输出该数字,否则输出-1
func duplicate(input []int) int {
	length := len(input)
	if length == 0 {
		return -1
	}
	for i := range input {
		/*
			6 3 1 0 2 5 3
			3 3 1 0 2 5 6
			0 3 1 3 2 5 6
			0
		*/

		for iter := 0; input[i] != i; iter++ {
			if iter >= length {
				break
			}
			tmp := input[input[i]]
			if input[i] == tmp {
				return tmp
			}
			input[input[i]] = input[i]
			input[i] = tmp
			//fmt.Println(input)
		}
		length--
	}
	return -1

}

func main() {
	testSlice := []int{6, 3, 1, 0, 2, 5, 3}
	fmt.Println(testSlice)
	fmt.Println(duplicate(testSlice))
}
