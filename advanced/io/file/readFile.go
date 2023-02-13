package main

import (
	"fmt"
	"os"
)

func ReadAtFile(path string, offset int64) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	buf := make([]byte, 1024) // 每次读取1024个内容
	fmt.Println("下面是文件内容:")
	_, _ = file.ReadAt(buf, int64(offset))
	fmt.Println(string(buf))
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

// ReadFileOffset 从指定位置读取文件
func ReadFileOffset(path string, offset int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	buf := make([]byte, 1024) // 每次读取1024个内容
	fmt.Println("下面是文件内容:")
	file.Seek(int64(offset), 0)
	for {
		fileLen, _ := file.Read(buf) // Read方法会改变文件当前的offset
		if fileLen == 0 {
			break
		}
		fmt.Println(fileLen, string(buf[:fileLen]))
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

// ReadFile 读取文件
func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024) // 每次读取1024个内容
	fmt.Println("下面是文件内容:")
	for {
		fileLen, _ := file.Read(buf) // Read方法会改变文件当前的offset
		if fileLen == 0 {
			break
		}
		fmt.Println(fileLen, string(buf[:fileLen]))
	}
	file.Close()
}

func main() {
	//ReadFile("1.txt")
	ReadFileOffset("2.txt", 3)
	ReadAtFile("2.txt", 3)
}
