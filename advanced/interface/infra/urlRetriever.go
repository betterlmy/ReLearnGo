package infra

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Retriever struct {
}

// Get 给定url保存到本地html文件
func (Retriever) Get(url string) error {
	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Println("get错误")
		return getErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("close错误")
			return // 退出当前函数
		}
	}(resp.Body) // 推迟执行关闭函数

	if resp.StatusCode != 200 {
		log.Println("状态码错误")
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	bytes, readErr := io.ReadAll(resp.Body) // 读取所有的内容
	if readErr != nil {
		log.Println("read错误")
		return readErr
	}
	// 将输出保存为html文件
	writeErr := os.WriteFile("download.html", bytes, 0666)
	if writeErr != nil {
		log.Println("write错误")
		return writeErr
	}

	return nil
}
