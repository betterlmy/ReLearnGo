package main

import (
	"fmt"
	x "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin 开始运行")
	r := x.Default()
	r.GET("/", func(ctx *x.Context) {
		ctx.String(200, "Zane's First Go Web (gin) Demo")
	})
	r.Run()
}
