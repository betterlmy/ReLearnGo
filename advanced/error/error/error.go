package main

import (
	"fmt"
	"math"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("Error!不能对负数`%v`进行开根号操作", float64(e))
}

// sqrt 安全的开根号处理
func sqrt(x float64) (float64, error) {
	if x < .0 {
		return 0, ErrNegSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	x := [2]float64{2.0, -1.0}
	for _, v := range x {
		res, err := sqrt(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("v开根号的结果是:", res)
		}

	}
}
