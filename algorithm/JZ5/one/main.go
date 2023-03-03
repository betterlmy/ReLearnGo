package main

import "fmt"

/*
请实现一个函数，将一个字符串s中的每个空格替换成“%20”。
例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
*/

func replaceSpace(s string) string {
	// write code here

	var out []rune
	if len(s) == 0 {
		return string(out)
	}
	for _, v := range s {
		if v != 32 {
			out = append(out, v)
		} else {
			out = append(out, []rune{37, 50, 48}...)
		}

	}
	return string(out)
}

func main() {
	s := "156 1"
	fmt.Print(replaceSpace(s))
}
