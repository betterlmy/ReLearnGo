## 接口(Interface)

### 示例代码

```go
type Traversal interface{
Traverse()
}
func main(){
traversal := getTraversal()
traversal.Traverse()
}
```

###  接口的概念
