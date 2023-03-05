package main

import "fmt"

// hash法
func duplicate1(numbers []int) int {
	// write code here
	exist := make([]bool, len(numbers))
	for _, v := range numbers {
		if exist[v] == true {
			return v
		} else {
			exist[v] = true
		}
	}
	return -1

}

func main() {
	testSlice := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(duplicate1(testSlice))
}
