package main

import "fmt"

func run(args ...string) {
	// do something
	fmt.Println(args)
}

func main() {
	go run("this is a new thread")
}
