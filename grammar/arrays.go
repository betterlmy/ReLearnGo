package main

import "fmt"

// arrays 定义数组的方法
func arrays() ([5]int, [3]int, [5]int) {

	var arr1 [5]int         // 默认为0
	arr2 := [3]int{1, 2, 3} //需要赋初值
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2)
	// 数组的遍历

	for _, i2 := range arr3 {
		fmt.Println(i2)
	}
	for i := range arr2 {
		fmt.Println(i)
	}
	return arr1, arr2, arr3
}

func printArray(arr [5]int) {
	fmt.Println("要求数组长度必须是5,否则无法传入")
	for i := range arr {
		fmt.Println(i)
	}
}
