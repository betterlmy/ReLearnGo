package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("1.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(file)
		file.Close()
	}
}
