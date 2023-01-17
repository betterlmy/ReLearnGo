# 基础语法部分

## 关键字

Go语言中共有`25`个关键字，如下所示：

```go
break       default     func interface select
case defer go map         struct
chan        else goto package switch
const fallthrough if range type
continue    for import return var
```

此外,包含`36`个预定义标识符,如下所示:

```go
Constants: true false iota nil
Types: int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error
Functions: make len cap new append copy close delete
complex real imag
panic recover
```

## 变量

[variable()函数](./main.go)
![](assets/OzmOsy.png)
![](assets/zuoIVz.png)

### 变量声明

Go语言中,变量需要声明后才可以使用,同一作用域内不支持重复声明,且声明后必须使用,否则会报错.  
Go语言推荐驼峰式命名法,例如: `userName`  
声明代码如下:

```go
//标准单个声明
var 变量名 变量类型
// 例如
var name string
var age int
var isOk bool

//批量声明
var (
name string
age int
isOk bool
)

//短变量声明,可以自动判别类型,但是无法在函数外使用此声明方法
age, time, des := 5, 7, "这是一个描述" 
```

如在函数外声明,只能使用标准声明方法,且并全局变量,而是包内变量

### 内建变量类型

* bool,string
* (u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr(指针)
* byte(8位),rune(字符,长度32位)
* float32,float64,complex64,complex128(复数)

### 强制类型转换

Go语言类型转换是强制的,例如说
sqrt()函数返回一个float64类型,但是我们需要一个int类型,那么就需要强制转换

```go
var c int
c = int(math.Sqrt(4))
//c = math.Sqrt(4)是会报错的
```

### 输出变量类型
    
```go
fmt.Printf("%T", c)
fmt.Println(reflect.TypeOf(c))
```
```