package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

type Rect struct {
}

// GetArea 求面积
func (r *Rect) GetArea(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// GetPerimeter 求周长
func (r *Rect) GetPerimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {

	// 1. new矩形对象
	rect := new(Rect)

	// 2. 将rect注册为一个rpc服务
	rpc.Register(rect)

	// 3. 绑定http
	rpc.HandleHTTP()
	fmt.Println("开始监听:19999")

	err := http.ListenAndServe(":19999", nil)
	if err != nil {
		log.Panic()
	}

}
