package main

import (
	"fmt"
	"os"
)

// deleteEmptyDir deletes an empty directory.
func deleteEmptyDir(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				deleteEmptyDir(dirPath + "/" + file.Name())
			}
		}
	}
	files, err = os.ReadDir(dirPath)
	if len(files) == 0 {
		os.Remove(dirPath)
		fmt.Println(dirPath, "已被删除")
	}

}

func main() {
	deleteEmptyDir("/Users/zane/Documents/GitHub/ReLearnGo/advanced/io/dir/testDir2")
}
