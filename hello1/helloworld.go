package hello

import (
	"fmt"
	"hello/hello2"
)

func Hello() {
	var str string = "helloworld"
	fmt.Print(str)
	fmt.Printf("%s", str)
	hello2.Hello2()
}
