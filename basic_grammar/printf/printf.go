package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var x = 2588
	var li = Person{"李梦洋", 18}
	zhang := Person{"小李2", 328}
	fmt.Printf("%#v \n", li)
	fmt.Printf("%+v \n", zhang)
	fmt.Printf("%v \n", zhang)
	fmt.Printf("%T \n", zhang)

	fmt.Printf("八进制:%o,十进制%d,十六进制:%X \n", x, x, x)
}
