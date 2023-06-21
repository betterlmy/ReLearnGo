package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	// 1. 建立与远程rpc的连接

	conn, err := rpc.DialHTTP("tcp", ":19999")
	if err != nil {
		log.Fatal("tcp连接失败")
	}

	// 2. 调用rpc方法
	ret := 0
	err2 := conn.Call("Rect.GetArea", Params{
		Width:  3,
		Height: 10,
	}, &ret)
	if err2 != nil {
		fmt.Println("求面积错误")
		log.Fatal(err2)
		return
	}
	fmt.Println("面积是:", ret)

	err2 = conn.Call("Rect.GetPerimeter", Params{
		Width:  3,
		Height: 10,
	}, &ret)
	if err2 != nil {
		fmt.Println("求周长错误")
		log.Fatal(err2)
		return
	}
	fmt.Println("周长是:", ret)

}
