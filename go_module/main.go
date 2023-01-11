package main

// 同一文件夹不允许有多个包
import (
	"fmt"
	xx "go_module/x" // 调用相同文件夹下的包需要添加package
	zx "go_module/z" // 因为x文件夹和z文件夹包含相同的包x,所以需要分别重命名
	"hello"          // 调用区别与当前工作目录的hello包,位于上层目录的hello1文件夹中,所以需要修改go.mod文件(添加require并使用repleace替换为本地路径)
)

func main() {
	fmt.Println("本地调包测试")
	xx.New()      // 调用当前工作目录下x文件夹中x包的New方法
	zx.New2()     // 调用当前工作目录下z文件夹中x包的New2方法
	zx.New()      // 调用当前工作目录下z文件夹中x包的New方法
	hello.Hello() // 调用其他工作目录下hello1文件夹中的Hello方法
}
