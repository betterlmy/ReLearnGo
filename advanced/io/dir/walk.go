package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkDir1(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	WalkDir1("/Users/zane/Documents/GitHub/ReLearnGo/advanced/")
}
