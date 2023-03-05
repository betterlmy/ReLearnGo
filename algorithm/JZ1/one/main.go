package main

import (
	"fmt"
)

func divide(a int, b int) int {
	//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31, 2^31−1]。本题中，如果除法结果溢出，则返回 2^31 − 1
	//min := int(-1 * math.Pow(2, 31))
	min := -2147483648
	//max := int(math.Pow(2, 31) - 1)
	max := 2147483647
	// 特殊情况

	if a == 0 {
		return 0
	}

	if b == 1 {
		if a >= max || a < min {
			return max
		}
		if a == min {
			return min
		}
		return a
	}

	if b == -1 {
		if a <= min || a > max+1 {
			return max
		}
		if a == max+1 {
			return min
		}
		return -1 * a
	}

	// 符号计算
	pos := 1
	if a < 0 {
		a = -a
		pos = -1 * pos
	}
	if b < 0 {
		b = -b
		pos = -1 * pos
	}

	result := 0
	for a-b >= 0 {
		a -= b
		result++
	}
	result *= pos

	return result

}
func main() {
	a, b := 10, 5

	fmt.Println(divide(a, b))

}
