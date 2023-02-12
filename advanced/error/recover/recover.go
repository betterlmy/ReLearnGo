package main

import "fmt"

// protect 以安全模式来运行所有传入的方法
func protect(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	f() // 方法f()会在宕机恢复后进行调用
}

func main() {
	protect(func() {
		fmt.Println("A")
		panic("A--panic")
	})

	protect(func() {
		fmt.Println("B")
		panic("B--panic")
	})
	fmt.Println("last line")
}
