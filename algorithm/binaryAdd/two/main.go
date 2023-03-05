package main

import (
	"fmt"
	"math"
)

/*
转 将string转为二进制,再转成十进制 相加后 在转成二进制,
*/

func addBinary(a string, b string) string {
	intA := str2Dec(a)
	intB := str2Dec(b)
	sum := intA + intB
	return Dec2Binstring(sum)
}

// Dec2Binstring 十进制转二进制字符串
func Dec2Binstring(x int64) string {
	result := make([]byte, 0)

	for {
		if x == 1 {
			result = append(result, '1')
			break
		} else if x == 0 {
			result = append(result, '0')
			break
		}
		if x%2 == 0 {
			result = append(result, '0')
		}
		if x%2 == 1 {
			result = append(result, '1')
		}
		x /= 2
	}
	result = reverse(result)
	return string(result)
}

// str2Dec 将str类型的二进制转为十进制int型
func str2Dec(str string) int64 {
	byteStr := reverse([]byte(str))
	res := 0.0
	for i, v := range byteStr {
		if v == '1' {
			res += math.Pow(2.0, float64(i))
		}
	}
	return int64(res)
}

func reverse(x []byte) []byte {

	length := len(x)
	if length == 0 {
		return nil
	}
	for i := 0; i < length/2; i++ {
		x[i], x[length-i-1] = x[length-i-1], x[i]
	}
	return x

}

func main() {
	a := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Println(addBinary(a, b))

}
