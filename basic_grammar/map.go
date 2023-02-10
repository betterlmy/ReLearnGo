package main

import (
	"fmt"
	"reflect"
)

func mapInit() {
	m := map[string]string{
		"name":  "李梦洋",
		"age":   "18",
		"email": "betterlmy@icloud.com", // 必须有逗号
	}

	m2 := make(map[string]int) // m2 空map
	var m3 map[string]int      // m3 nil

	fmt.Println(m, m2, m3)
	fmt.Println(m["age1"])

	m["qq"] = "283738217"
	// map 的遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
	x, isExist := m["gender"]
	fmt.Println(reflect.TypeOf(x), isExist)

	// 删除元素
	delete(m, "qq")
	length := len(m)
	fmt.Println(length)
}
