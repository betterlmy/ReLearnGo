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

## 变量与常量

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

### 常量的定义

```go
const (
filename = "abc.txt"
a, b = 3, 4
)
var c int
c = int(math.Sqrt(a*a + b*b))
fmt.Println(filename, c)
```

const 数值可以作为各种类型使用,例如上述定义ab为int类型,但是在Sqrt函数中,要求参数为float64类型,仍然可以使用.

### 使用常量定义枚举类型

```go
//普通枚举类型
const(
cpp = iota
java
python
golang
js
)

//自增枚举类型
const (
b = 1 << (10 * iota)
kb
mb
gb
tb
pb
)

```

### 回顾

* 变量类型写在变量名之后
* 编译器可推测变量类型
* 有char，只有rune
* 原生支持复数类型

## 条件语句

### if语句

```go
func bounded(v int) int {
if v >= 100 {
return 100
} else if v < 0 {
return 0
} else {
return v
}
}
```

使用[io操作](io/io.go)进行条件语句示例.  
if语句可以直接赋值,且同时变量作用域只限制在这个if语句内.

### switch语句

switch会自动break,除非使用fallthrough强制执行后面的case代码.

```go

// eval switch有表达式的情况
func eval(a, b int, op string) float64 {
var result float64
switch op {
case "+":
result = float64(a + b)
case "-":
result = float64(a - b)
case "*":
result = float64(a * b)
case "/":
if b == 0 {
panic("除数不能为0")
}
result = float64(a) / float64(b)
default:
panic("不支持的操作子:" + op)
}
return result
}

// grade switch没有表达式的情况,需要添加判断语句
func grade(score int) string {
g := ""
switch {
case score < 0 || score > 100:
panic(fmt.Sprintf("Wrong score: %d", score))
case score < 60:
g = "F"
case score < 80:
g = "C"
case score < 90:
g = "B"
case score <= 100:
g = "A"
default:
panic(fmt.Sprintf("Wrong score: %d", score))
}
return "得分为" + g
}
```

## for循环

### 初始化,条件,后置语句三要素都有

```go
var sum int
for i := 0; i < 10; i++ {
fmt.Println(i)
sum += i
}
```

### 单条件循环(类while)

```go
func PrintFile(filename string) {
// 读取文件内容
contents, err := os.Open(filename)
if err != nil {
panic(err)
}
scanner := bufio.NewScanner(contents)
for scanner.Scan() {
fmt.Println(scanner.Text())
}
}
```

### while True循环

```go
for {
fmt.Println("1")
}
```

## 函数

* 函数没有默认参数
* 函数可以返回多个值
* 函数返回多个值时,可以起名字,对调用者而言没有区别

```go
func div(a, b int) (q, r int) {
q = a / b
r = a % b
return
}
```

* 函数可以作为变量的值

```go
func apply(op func (float64, float64) float64, a, b float64) float64 {
p := reflect.ValueOf(op).Pointer() //返回指向函数的指针
opName := runtime.FuncForPC(p).Name() //返回函数名
fmt.Println("执行函数:", opName, "with args", a, " ", b)
return op(a, b)
}
//在使用时可以直接使用 
apply(func (a, b float64) float64 {return a+b}, 1, 2) 达到使用匿名函数的效果
```

* 可变参数的设置

```go
func sum(num ...int) int {
s := 0
for _, j := range num {
s += j
}
return s
}
```

### 指针
* 指针不能做运算
* 指针的值为nil时,不能进行解引用操作
* Go语言的参数传递为引用传递
```go
func passByValue(a int) {
	a += 1
}

func passByPointer(a *int) {
	*a += 1
}

func pointerTypeCheck() {
	a, b := 3, 3
	passByValue(a)
	passByPointer(&b)
	fmt.Println("值传递", a)
	fmt.Println("指针传递", b)

}
```

## 数组
### 数组的定义方法
```go
	var arr1 [5]int         // 默认为0
	arr2 := [3]int{1, 2, 3} //需要赋初值
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
```

### 数组是值类型
```go
func printArray(arr [5]int) {
	fmt.Println("要求数组长度必须是5,否则无法传入")
	for i := range arr {
		fmt.Println(i)
	}
}
    arr1, arr2, arr3 := arrays()
    printArray(arr1)
    printArray(arr2) // 无法使用 会报错,因为数组是值类型,传递的是值的拷贝
    printArray(arr3)
```