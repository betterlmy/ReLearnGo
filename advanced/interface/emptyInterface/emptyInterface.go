package main

import "fmt"

func Log(i interface{}) {
	fmt.Printf("type:%T,value:%v\n", i, i)
}

func main() {

	x := 1
	var interfs []interface{} = []interface{}{
		x,
		"hello",
		true,
		&x,
		struct { // 匿名结构
			Name string
		}{"小明"},
		[]int{1, 2, 3},
	}
	for _, interf := range interfs {
		Log(interf)
	}
	str := interfs[1] //"hello",此时str类型为空接口
	// var d string = str // 不能直接通过类型转换,因为变量str类型是空接口,不能直接转换为string类型
	d := str.(string) //类型断言,将空接口转换为string类型
	fmt.Println(d)
}
