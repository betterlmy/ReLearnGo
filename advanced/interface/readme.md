## 接口(Interface)

### 概念

接口是Go语言中一种类型，表示`一组方法签名（方法名、参数、返回值列表）的集合`。
接口类型不包含具体地实现，但是任何类型（包括结构体、函数、别名类型等）如果实现了接口中的所有方法，就可以被认为是该接口的实例。
`接口由使用者进行定义`,即接口由测试等人进行定义使用

### 示例代码

```go
type retriever interface {
Get(string) error // 这个接口包含着Get这个方法,同时限定了方法的参数和返回值
}
```

上述代码定义了一个接口retriever,接口中限定了一个方法Get,这个方法的参数是string类型,返回值是error类型.

```go
type Retriever1 struct {
}
func (Retriever1) Get(url string) error {
    return nil
}

type Retriever2 struct{}

func (Retriever2) Get(url string) error {
fmt.Println("Get测试")
return nil
}

```

上述代码定义了两个结构,分别是Retriever1和Retriever2,这两个结构都实现了retriever接口中的Get方法,因此这两个结构都可以被认为是retriever接口的实例.可以使用以下方式进行使用:

```go
var retriever1 retriever = infra.Retriever{}
var retriever2 retriever = testing.Retriever{}

err1 := retriever1.Get(url)
err2 := retriever2.Get(url)
```

上述代码将两个不同的结构实例赋值给了接口,并且可以直接使用Get方法,虽然名字相同,但是实现的功能由各自的结构体决定.

### 功能意义

接口在Go语言中起到了非常重要的作用：

* 实现多态：接口允许不同类型的值被统一地处理，因为它们都实现了同一个接口。
* 封装实现细节：接口的实现可以被隐藏，从而保护其实现细节不被外界误用或滥用。
* 增强代码灵活性：接口允许不同类型的实现在不同情况下被动态地切换。

总之，接口是Go语言中一种非常重要的抽象概念，能够帮助我们提高代码的灵活性、可扩展性和可维护性。

## Duck Typing

![小黄鸭](../assets/IJBTNC.png)

### 理解概念

一个对象像鸭子一样走路,像鸭子一样叫,那么这个对象就可以是鸭子.
duck typing是编程语言中的一个概念,意味着`像鸭子一样看待对象`

### 基本思想

在使用对象时,不用关心它的类型具体是什么,而是关心它能够支持哪些属性和方法.如果一个对象满足我们的需求,那么我们就可以使用它.
Duck typing与静态类型语言中的类型系统相比，更加灵活和弹性。这种方法主要用于动态语言（如 Python），也在Go语言中被广泛使用（通过接口实现）。

### 示例代码

#### python

```python
def download(x):
    return x.get("google.com")
```

Python实现duck typing的代码如上,这段代码中的x可以是任何类型的对象, 且不对x进行验证,执行才知道x有没有get方法

#### go

```go
type Retriever interface {
Get(url string) string
}
```

在Go中,duck typing的实现是通过接口实现的,上述代码定义了一个接口retriever,接口中限定了一个方法Get,这个方法的参数是string类型,返回值是string类型.
一个结构体需要`定义Get方法,并且满足Get方法的参数和返回值类型`,那么这个结构体才可以被认为是retriever接口的实例.

### Go语言接口的优点

* 同时具有python,c++的duck typing的灵活性
* 又具有java的类型检查

