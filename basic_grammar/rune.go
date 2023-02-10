package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func runeInit() {
	s := "Yes我爱你!" // 一个中文占4个字符长度,即32位,在go中 rune=int32,所以说string就是由rune组成的
	fmt.Println("直接输出len(s),长度为:13", len(s))
	for _, v := range s {
		fmt.Print(string(v) + " ")

		fmt.Print("十进制:" + strconv.Itoa(int(v)))
		fmt.Printf(" %T ", v)
		fmt.Printf("16进制:%X \n", v)
	}
	fmt.Println("使用for循环,循环次数为:7")
	// 为了获取rune的长度,需要使用utf-8库
	fmt.Println("使用utf-8库,获取的长度", utf8.RuneCountInString(s))
	fmt.Println('a', "a") //单引号为utf-8的rune类型,双引号为string类型
	a := 'a'
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("%c", a) // 使用string()方法或者%c将单个rune转换为string
}
