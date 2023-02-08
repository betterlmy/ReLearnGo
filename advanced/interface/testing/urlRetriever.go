package testing

import "fmt"

type Retriever struct{}

func (Retriever) Get(url string) error {
	fmt.Println("Get测试")
	return nil
}
