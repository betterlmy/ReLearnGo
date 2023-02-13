package main

import (
	"fmt"
	"os"
)

func createOneDir(path string, dirName string) {
	if path == "" {
		path, _ = os.Getwd()
	}
	newDirPath := path + "/" + dirName
	err := os.Mkdir(newDirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		os.Chmod(newDirPath, 0777)
		fmt.Println("新建目录:", newDirPath)
	}

}

func createAllDir(path string, dirName string) {
	if path == "" {
		path, _ = os.Getwd()
	}
	newDirPath := path + "/" + dirName
	err := os.MkdirAll(newDirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		os.Chmod(newDirPath, 0777)
		fmt.Println("新建目录:", newDirPath)
	}

}
func main() {
	createOneDir("/Users/zane/Documents/GitHub/ReLearnGo/advanced/io/dir", "testDir1")
	createAllDir("/Users/zane/Documents/GitHub/ReLearnGo/advanced/io/dir/testDir2", "testDir3/testDir4")
}
