package main

import (
	"fmt"

	x "github.com/gin-gonic/gin"
)

func main() {
	start()
}

func start() {
	fmt.Println("Gin 开始运行")
	r := x.Default()
	str := "Zane's First Go Web (gin) Demo"
	r.GET("/", func(ctx *x.Context) {
		ctx.String(200, str)
	})
	err := r.Run()
	if err != nil {
		return
	}
	fmt.Println()
}
