## 代码问题

### 9. 写出打印结果

```go
type People struct {
	name string `json:"name"`
}

func main() {
	js := `{
"name":"11" 
}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
```

上面的代码中,js字符串中name是非导出属性,且在People这个结构体中,people也是非导出的,所以说使用json.Unmarshal无法将js字符串转换为结构体对象.

### 10. 下面代码是有问题的

```go
type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}
func main() {
	p := &People{}
	p.String()
}

```

在go中,string类型其实已经实现了String()方法,fmt.Sprintf("print: %v", p)这段代码会调用p.String()也就是说产生了循环调用(递归?).

### 11. 找出下面代码的问题