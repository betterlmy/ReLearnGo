package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return -1
	}
	/* 	因为要求空间是O(1),所以只能有一个变量
	我们用一个变量times 来计算一个范围内数字出现的次数,比如果我计算出1-5这五个数字出现的次数times>5,
	那么证明1-5之间一定会有重复
	*/
	return getRepeatNumber(nums, 0, n-1)
}

func getRepeatNumber(nums []int, min, max int) int {
	length := max - min + 1
	if length == 1 {
		return max
	}
	mid := min + length/2
	times := 0
	for _, num := range nums {
		for i := min; i <= mid; i++ {
			if num == i {
				times++
				break
			}
		}
	}
	leftLength := mid - min + 1
	fmt.Printf("Checking [%d~%d],times:%d\n", min, mid, times)

	if times > leftLength {
		// mid-min+1指的是当前左侧要检测的数字数量,当这些数字出现的数量大于mid-min+1时,说明重复出现的地方一定在这个侧边
		if leftLength == 2 {
			minTimes := 0
			midTimes := 0
			for _, num := range nums {
				if num == min {
					minTimes++
					if minTimes >= 2 {
						return min
					}
				} else if num == mid {
					midTimes++
					if midTimes >= 2 {
						return mid
					}
				}
			}
		}
		return getRepeatNumber(nums, min, mid)
	}

	times = 0
	for _, num := range nums {
		for i := mid + 1; i <= max; i++ {
			if num == i {
				times++
				break
			}
		}
	}
	fmt.Printf("Checking [%d~%d],times:%d\n", mid+1, max, times)

	rightLength := max - mid
	if times > rightLength {
		if rightLength == 2 {
			midTimes := 0
			maxTimes := 0
			for _, num := range nums {
				if num == mid+1 {
					midTimes++
					if midTimes >= 2 {
						return num
					}
				} else if num == max {
					maxTimes++
					if maxTimes >= 2 {
						return num
					}
				}
			}

		}
		return getRepeatNumber(nums, mid+1, max)
	}
	return -1
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 12, 9, 10, 11, 12}
	fmt.Println(slice)
	fmt.Println(findRepeatNumber(slice))
}
