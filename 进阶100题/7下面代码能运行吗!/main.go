package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	// 此时s中的Param的值是nil
	s.Param = make(map[string]interface{}) // 需要初始化
	s.Param["RMB"] = 100
	fmt.Println(s.Param)
}
