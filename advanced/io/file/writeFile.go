package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("3.txt") // 如果文件已经存在 清空文件
	if err != nil {
		panic(err)
	}
	data := "我是一行数据\r\n"
	for i := 0; i < 3; i++ {
		write, err := file.Write([]byte(data))
		if err != nil {
			return
		}
		fmt.Println("新写入了", write, "字节")
	}
	file.WriteAt([]byte("我是一行数据"), int64(3*len(data)+5))
	file.Close()
}
