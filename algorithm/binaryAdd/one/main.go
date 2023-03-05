package main

import "fmt"

// 进制法 按位相加
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a // 保证 len(a)>=len(b)
	}

	//翻转方便计算
	byteA := reverse([]byte(a))
	byteB := reverse([]byte(b))

	lenA := len(a)
	lenB := len(b)

	result := make([]byte, lenA+1) // +1 是因为结果最多增加一位

	c := 0 // 进位
	for i := 0; i < lenA; i++ {
		if byteA[i] == '1' {
			c++
		}
		if i < lenB && byteB[i] == '1' {
			// 超出长度按0计算 所以不会有进位
			c++
		}
		if c == 1 || c == 3 {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
		c >>= 1 //右移一位 如果是3 则为1 表明有一个进位
		// 或者可以 c/=2

	}
	if c == 1 {
		result[lenA] = '1'
	}
	result = reverse(result)

	if c == 1 {
		// 如果最后c=1 说明延长了结果 需要全部输出
		return string(result)
	} else {
		// 否则需要把多补的0删掉
		return string(result[1:])
	}

}

func reverse(x []byte) []byte {
	// 翻转 从小往大进行计算
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-i-1] = x[n-i-1], x[i]
	}
	return x
}

func main() {
	fmt.Println(addBinary("11", "10"))

}
