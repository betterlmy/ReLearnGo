package main

import (
	"errors"
	"fmt"
	"runtime"
)

func main() {
	err := errors.New("这是一个新的错误")
	var err2 error
	fmt.Println(err)  // 这是一个新的错误
	fmt.Println(err2) // <nil>

	if _, _, line, ok := runtime.Caller(0); ok == true { //通过调用runtime包获取当前代码所在的行数
		err := fmt.Errorf("error in line %d,错误发生在第%d行", line, line)
		fmt.Println(err)
	}

}
