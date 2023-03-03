package main

import "fmt"

/*
 * 10ms4932kb
 */

func Find(target int, array [][]int) bool {
	// write code here

	/*
		在一列中,比他大的只有向下,比他小的话 只有向上
		在一行中,比他大的只有向右,比他小只有向左
		假设 target=7:
		从左上角开始,左上角=1,此时我们可以向右,也可以向下,不确定	左上角=10,直接输出不存在
		从右上角开始,右上角=1,此时我们只可以向下  				右上角=10,只需要向左
		从左下角开始,左下角=1,此时只用往右寻找					左下角=10,只用网上寻找
		从右下角开始,右下角=1 直接输出不存在						右下角=10,我们可以向上,也可以向左,不确定
		因此我们可以从右上角开始或者左下角开始
		此算法O(m+n)


	*/
	// 从右上角开始
	width := len(array[0]) - 1
	height := len(array) - 1
	if width < 0 {
		return false
	}
	n := 0
	for width >= 0 && n <= height {
		fmt.Println("当前正在寻找", n, width)
		if array[n][width] > target {
			width--
		} else if array[n][width] < target {
			n++
		} else {
			return true
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
		{6, 15, 18, 20},
	}
	target1 := 7

	target2 := 3

	test(array1, target1)
	test(array1, target2)
}
