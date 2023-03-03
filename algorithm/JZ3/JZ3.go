package main

import "fmt"

var exist []int

func duplicate(numbers []int) int {
	// write code here
	// 在一个长度为n的数组中，所有的数字都在[0,n-1]的范围内，数组中的某些数字是重复的，但不知道有几个是重复的，也不知道每个数字重复几次，找出数组中任意一个重复的数字，
	//[2,3,1,0,2,5,3]
	for _, v := range numbers {
		if checkDuplicate(v) {
			return v
		}
	}
	return -1

}

func checkDuplicate(x int) bool {
	for _, v := range exist {
		if x == v {
			return true
		}
	}
	exist = append(exist, x)

	return false
}

func main() {
	testSlice := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(duplicate(testSlice))
}
