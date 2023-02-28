package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 要匹配的字符串
	// 判断字符中是test
	// 输出正则表达式的字符串表示形式
	fmt.Println("正则表达式的字符串表示形式为:")
	s := "test"

	// 定义正则表达式

	// 使用正则表达式进行匹配
	if ok, _ := regexp.MatchString("^((?!test).)*$", s); ok == true {
		fmt.Println("字符串中不包含test")
	} else {
		fmt.Println("字符串中包含test")
	}
}
