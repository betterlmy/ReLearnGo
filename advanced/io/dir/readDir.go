package main

import (
	"fmt"
	"os"
)

func ListDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, filename := range files {
		if filename.IsDir() {
			ListDir(dir + filename.Name() + "/")
		}
		fmt.Printf("文件名:%s%s\n", dir, filename.Name())
	}

	return nil
}

func main() {
	dir := "/Users/zane/Documents/GitHub/ReLearnGo/advanced/"
	ListDir(dir)
}
