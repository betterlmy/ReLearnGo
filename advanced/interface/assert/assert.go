package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// TypeCheck 判断接口实例的类型
func TypeCheck(v interface{}) bool {
	switch v.(type) {
	case int:
		fmt.Println(v, "是int类型的")
	case string:
		fmt.Println(v, "是string类型的")
	case Person:
		fmt.Println(v, "是Person类型的")
	case float64:
		fmt.Println(v, "是float64类型的")
	default:
		fmt.Println(v, "是未知类型的")
		return false
	}
	return true
}

func main() {
	var s = []interface{}{1, "hello", Person{"Tom", 20}, 3.14, true}
	for _, v := range s {
		TypeCheck(v)
	}
}
