go语言高级编程

## 接口

### Duck Typing

![小黄鸭](assets/IJBTNC.png)

#### 理解概念

一个对象像鸭子一样走路,像鸭子一样叫,那么这个对象就可以是鸭子.
duck typing是编程语言中的一个概念,意味着`像鸭子一样看待对象`

#### 基本思想

在使用对象时,不用关心它的类型具体是什么,而是关心它能够支持哪些属性和方法.如果一个对象满足我们的需求,那么我们就可以使用它.
Duck typing与静态类型语言中的类型系统相比，更加灵活和弹性。这种方法主要用于动态语言（如 Python），也在Go语言中被广泛使用（通过接口实现）。

#### 示例代码

```python
def download(x):
    return x.get("google.com")
```

Python实现duck typing的代码如上,这段代码中的x可以是任何类型的对象, 且不对x进行验证,执行才知道x有没有get方法

```go
type Retriever interface {
Get(url string) string
}
```

在Go中,duck typing的实现是通过接口实现的,上述代码定义了一个接口retriever,接口中限定了一个方法Get,这个方法的参数是string类型,返回值是string类型.
一个结构体需要`定义Get方法,并且满足Get方法的参数和返回值类型`,那么这个结构体才可以被认为是retriever接口的实例.

### 概念

接口是Go语言中的一种类型，表示`一组方法签名（方法名、参数、返回值列表）的集合`。
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
[完整代码](interface/downloader.go)

另有通过[数据库连接](interface/sql/main.go)演示了面向过程和面向对象(多态)的区别

Go语言中,一个类型可以实现多个接口,当然,多个类型也可以实现同一个接口,任何一个类型必然实现了空接口,因此空接口可以表示任何类型.

### 接口赋值

接口赋值分为两种情况:

- 接口变量赋值给接口变量
- 满足接口方法的自定义类型的实例赋值给接口变量

如果自定义类型实现了某个接口类型所声明的方法,那么这么自定义类型的值就可以赋值给这个接口,这个赋值会把用户定义的类型的值存入接口类型的值,即类型的转换.

例如上面数据库连接中,声明了一个Mysql类型的变量m,我们可以直接将m赋值给IDatabaser,这样m的类型不再是Mysql而是IDatabaser.

在接口变量中,一共存储了两部分的信息:

- 分配给接口变量的具体值
- 值的类型的描述器(源类型)

当一个接口I1实现了方法A()和方法B()时,而接口I2只实现了方法A(),那么接口I1的值可以赋值给接口I2,但是接口I2的值不能赋值给接口I1.
因为I2的所有方法对于I1接口的实例而言都是可执行的,反之则无法保证.

### 空接口

`空接口(interface{})`是Go语言中最特殊的接口,Go语言中所有的类都可以认为是继承于空接口,因此空接口可以表示任何类型.

#### 空接口定义

空接口不包含任何的方法,所以 所有的类型都实现了空接口,因此空接口可以存储任意类型的数值.例如

```go
package main

import "fmt"

func Log(i interface{}) {
	fmt.Printf("type:%T,value:%v\n", i, i)
}

func main() {

	x := 1
	var interfs []interface{} = []interface{}{
		x,
		"hello",
		true,
		&x,
		struct { // 匿名结构
			Name string
		}{"小明"},
		[]int{1, 2, 3},
	}
	for _, interf := range interfs {
		Log(interf)
	}
}

/* Output:
type:int,value:1
type:string,value:hello
type:bool,value:true
type:*int,value:0x14000126008
type:struct { Name string },value:{小明}
type:[]int,value:[1 2 3]
*/
```

上面代码展示了空接口可以存储任意类型的数值.

#### 空接口取值

空接口的值可以通过类型断言将原来的值赋值给其他原生结构,例如

```go
str := interfs[1] //"hello",此时str类型为空接口
// var d string = str // 不能直接通过类型转换,因为变量str类型是空接口,不能直接转换为string类型
d := str.(string) //类型断言,将空接口转换为string类型
fmt.Println(d)
```

#### 空接口的应用

空接口可以作为函数的参数,这样函数就可以接收任意类型的参数,例如fmt.Println()函数就是一个接收空接口的函数.

```go
func Println(a ...interface{}) (n int, err error)
```

### 类型断言

`类型断言(assert)`是使用在接口变量上的操作,简单来说,接口类型向普通类型的转换就是类型断言(
猜,猜对了返回正确的,猜错了就返回初值).
基本语法如下:

```go
t, ok := X.(T)
```

其中X是一个接口变量,T是一个类型,t为X转换为T类型后的变量,ok为一个bool类型,表示断言是否成功. 如果失败,t则为T类型的零值,ok为false.

接口的断言有两种方式:

- ok-pattern
- switch-type

#### ok-pattern

一般来说 如果要断言的种类较少时,使用ok-pattern,其结构如下:

```go
if value, ok := x.(T);ok{
// T为具体的类型
// do something
}
```

#### switch-type

当断言的接口类型较多时,使用switch-type,其结构如下:

```go
switch value := x.(type) {
case T1:
// do something
case T2:
// do something
case T3:
// do something
case T4:
// do something
case T5:
// do something
case T6:
// do something
}
```

[示例代码](interface/assert/assert.go)

### 功能意义

接口在Go语言中起到了非常重要的作用：

* 实现多态：接口允许不同类型的值被统一地处理，因为它们都实现了同一个接口。
* 封装实现细节：接口的实现可以被隐藏，从而保护其实现细节不被外界误用或滥用。
* 增强代码灵活性：接口允许不同类型的实现在不同情况下被动态地切换。

总之，接口是Go语言中一种非常重要的抽象概念，能够帮助我们提高代码的灵活性、可扩展性和可维护性。

### 拓展

Go语言的接口是非侵入式的.

## 错误处理

### [error](error)

Go语言中没有异常机制,但是有一个error接口,可以通过error接口来处理错误.

#### 自定义error

除了使用errors包中的New方法来创建自定义异常,还可以通过接口来创建.  
Go语言中的error类型实际上是一个接口,只要实现了Error() string方法即可,源代码如下:

```go
type error interface{
Error() string
}
```

自定义一个类型,实现返回string类型的Error()方法即可,代码如下:

```go
type ErrNegSqrt float64
func (e ErrNegSqrt) Error() string{
return fmt.Sprintf("不能对负数`%v`进行开根号", e)
}
```

[完整调用代码](error/error/error.go)

### [panic宕机](error/panic)

一般而言,只有当程序发生不可逆的错误时,才需要使用`panic`方法来触发宕机.  
panic方法是Go语言的内置函数,使用该方法后,程序将直接中断.
panic的源代码如下,可以看到panic方法接收一个interface{}类型的参数,该参数可以是任意类型.

```go
func panic(v interface{})
```

#### 调用panic的情形:

* 程序处于失控状态且无法恢复,继续执行会影响其他程序时
* 发生不可预知的错误时
* 程序运行到不应该运行到的地方时

### [recover宕机恢复](error/recover)

Go语言通过内置函数recover来捕获宕机,类似于其他语言的try-catch机制  
在使用panic方法触发宕机之后,且在退出当前函数前,会延迟调用defer语句

```go
package main

import "fmt"

func protect() {
	defer func() {
		fmt.Println("func protect 退出")
	}()
	panic("panic in protect")
}

func main() {
	defer func() {
		fmt.Println("func main 退出")
	}()
	protect()
	fmt.Printf("这段代码无法被执行")
}

//输出
//func protect 退出
//func main 退出
//panic: panic in protect

```

由于defer语句低延迟执行的特性,我们可以通过`defer语句+匿名函数+recover方法`来实现对宕机的捕获,代码如下:

```go
package main

import "fmt"

func protect() {
	defer func() {
		if err := recover(); err != nil { // 可以获取到panic传入的参数
			fmt.Println("recover from ", err)
		}
	}()
	panic("panic in protect")
}

func main() {
	protect()
	fmt.Printf("这段代码无法被执行")
	// 执行结果
	//recover from  panic in protect
	//这段代码无法被执行

}
```

上述代码中,在protect函数中,首先defer了一个匿名函数,这个函数主要用于panic抓取,当panic产生了宕机之后,并不会中断整个程序的运行,而是进入到了defer语句中,而defer语句中的匿名函数做到了panic的恢复工作recover,且recover可以获取panic传入的参数,所以说整个main()
函数足以全部执行

实际应用中,我们可以通过写入安全函数的方式,即使发生panic 主程序仍然能够继续正常运行,[示例代码](error/recover/recover.go)

## 文件处理

Go语言对I/O操作主要提供了下列几个官方标准库:

| 包名            | 作用                               |
|---------------|----------------------------------|
| os            | 提供了平台无关的接口,用于操作系统底层的文件,目录,进程,信号等 |
| io            | 为io原语(i/o primitives)提供基本的接口     |
| fmt           | 格式化io,实现对控制台的输入和输出               |
| bufio         | 实现带缓冲io                          |
| ~~io/ioutil~  | 自Go1.16之后已被删除                    |
| path/filepath | 实现了跨平台的文件路径操作,提供了对文件路径的处理方法      |

具体包内接口的用法可以参考[中文文档](https://studygolang.com/pkgdoc),[官方文档](https://pkg.go.dev/std)

### 目录操作

想到读取一个目录的内容,需要用到os库(或io/ioutil包)的ReadDir方法,该方法返回一个os.DirEntry,该类型包含了文件的基本信息,如文件名,文件大小,文件权限等

```go
func ReadDir(dirname string) ([]os.DirEntry, error)
```

[完整代码](io/dir/readDir.go)

#### 获取目录下所有的文件

手动遍历代码

```go
func ListDir(dir string) error {
files, err := os.ReadDir(dir)
if err != nil {
return err
}

for _, filename := range files {
if filename.IsDir() {
ListDir(dir + filename.Name() + "/")
}
fmt.Printf("文件名:%s%s\n", dir, filename.Name())
}

return nil
}

```

使用Walk函数

```go
func WalkDir(dir string) {
err := filepath.Walk(dir, func (path string, info os.FileInfo, err error) error {
if info == nil {
return err
}
if info.IsDir() {
return nil
}
println(path)
return nil
})
if err != nil {
fmt.Printf("filepath.Walk() returned %v\n", err)
}
}
```

#### 创建目录

[完整代码](io/dir/mkdir.go)  
使用os库中的`Mkdir`接口

```go
func Mkdir(name string, perm FileMode) error{

}
```

`Mkdir`接口使用指定的名称和权限创建一个目录,错误情形:

* 要新建的目录已经存在
* 父目录存在
* 子目录是多个需要新建的目录
* 权限不足

如果需要一次性新建一个多级目录,需要使用`MkdirAll`接口

#### 删除目录

删除目录主要使用到`os.Remove(path string)`方法  
如果目录非空,则会报错,如果需要删除非空目录,需要使用`os.RemoveAll(path string)`方法  
递归删除目录下所有的空文件夹,[完整代码](io/dir/rmdir.go)

### 文件操作

#### 文件权限

在linux中,文件的权限分为读r,写w和执行x.具体数值与作用见下表

| 权限  | 数值  | 作用                         |
|-----|-----|----------------------------|
| r   | 4   | 可以读取文件内容;浏览当前目录            |
| w   | 2   | 可以写入文件内容;删除或者移动当前目录内的目录与文件 |
| x   | 1   | 可以执行文件或进入目录                |

#### 文件的创建与打开

文件的创建与打开使用`os.OpenFile()`方法,该方法的定义如下:

```go
func OpenFile(name string, flag int, perm FileMode)(file *file, err error)
}
```

参数flag用于指定文件的访问模式,常见的值已经已经在os库中定义为了常量:

```go
const (
O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
O_CREATE int = syscall.O_CREAT // 如果不存在将创建一个新文件
O_EXCL   int = syscall.O_EXCL // 和O_CREATE配合使用,文件必须不存在
O_SYNC   int = syscall.O_SYNC // 打开文件用于同步I/O
O_TRUNC  int = syscall.O_TRUNC // 如果可能,打开时清空文件
)
```

其中`O_CREATE`和`O_EXCL`常常一起使用,用于确保文件的创建是安全的,`O_EXCL`要求文件必须不存在,否则会报错,`O_CREATE`
则是如果文件不存在,则创建文件,存在的话则不进行任何操作.  
`O_RDONLY`,`O_RDONLY`,`O_RDWR`应该只指定其中一个,剩下的通过或操作符`|`进行组合,如`O_CREATE|O_APPEND`
表示创建文件并且写入时将数据附加到文件尾部.  
该函数内部会给flags加上`syscall.O_CLOEXEC`，在fork子进程时会关闭通过OpenFile打开的文 件，即子进程不会重用该文件描述符。

参数`perm`指定了文件的模式和权限位,类型是`os.FileMode`,且同样在os中定义为常量

```go
const(
ModeDir FileMode = 1 << (32 - 1 - iota) // 目录
ModeAppend                              // a:只能写入,且写入时将数据附加到文件尾部
ModeExclusive                           // 文件用于同步I/O
ModeTemporary                           // 文件是临时文件
ModeSymlink   // 文件是符号链接
ModeDevice    // 文件是设备文件,一般是指系统的外部设备
ModeNamedPipe // 文件是命名管道
ModeSocket    // 文件是Unix域socket
ModeSetuid    // 执行此文件时具有文件所有者的用户ID
ModeSetgid     // 执行此文件时具有文件所有者的组ID
ModeCharDevice // 文件是字符设备,一般指系统的终端设备
ModeSticky     // 只有拥有者才能删除此文件
ModeIrregular  // 文件是不规则的,即既不是目录,也不是常规文件

ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice // 文件类型掩码
ModePerm FileMode = 0777 // Unix权限位掩码
)
```

`OpenFile`是一个更加底层的打开函数,更多人使用`Open`或者`Create`代替本函数.

```go
func Open(name string)(*File, error){
return OpenFile(name, O_RDONLY, 0)
}

func Create(name string)(*File, error){
return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666
```

#### 文件读取

`File`类型的`Read`方法用于读取文件内容,其定义如下:

```go
func (f *File) Read(b []byte) (n int, err error)
```

Read方法从文件中每次读取len(b)
字节数据并写入到byte数组中,他返回读取的字节数或额错误信息,如果读取到文件末尾,则返回0和`io.EOF`错误.  
[示例代码](io/file/readFile.go)

当遇到特别大的文件时,并且只需要读取文件最后部分的内容时,可以使用`Seek`方法将文件指针移动到文件末尾,然后再使用`Read`
方法读取文件内容.或者使用ReadAt方法,该方法可以指定读取的位置和长度.

```go
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
```

ReadAt方法从文件的指定位置读取指定长度的内容,如果读取到文件末尾,则返回0和`io.EOF`错误.
[示例代码](io/file/readFile.go)

#### 文件写入

`File`类型的`Write`方法用于向文件中写入内容,其定义如下:

```go
func (f *File) Write(b []byte) (n int, err error)
// 参数b为要写入的内容
// 返回值n为写入的字节数
```

#### 文件删除

`os`包中的`Remove`函数用于删除文件,其定义如下:

```go
func Remove(name string) error
```

`Remove`函数只能删除文件,不能删除目录,如果要删除目录,需要使用`RemoveAll`函数,其定义如下:

```go
func RemoveAll(path string) error
```

### Json处理

JSON(JavaScript Object Notation)是一种轻量级的数据交换格式,Json最初属于JS的一部分吗,后来由于良好的可读性和跨语言的特性,现已独立于语言,被广泛使用.
[Json官网](http://json.org/)
Json中的键都是字符串的形式,值可以取任意类型.他有以下三种结构:

```json
//值为字符串或数组类型
{
  "name": "Jhon",
  "age": 20
}

//JSON数组
[
  {
    "name": "Jhon",
    "age": 20
  },
  {
    "name": "Dan",
    "age": 21
  }
]

//值为对象类型
{
  "name": "Jhon",
  "age": 20,
  "address": {
    "city": "Beijing",
    "street": "Xidan"
  }
}
```

Go语言提供了`encoding/json`包用于处理Json数据.

#### Json编码

json包中的`Marshal`函数用于将Go语言中的数据类型转换为Json格式的数据,其定义如下:

```go
func Marshal(v interface{}) ([]byte, error)
```

为了让输出的Json字符串更加直观,可以使用`MarshalIndent`函数,其定义如下:

```go
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

demo

```json
func map2Json() {
  m: =
  make(map[
  string
]interface{}, 6)
m["name"] = "小王子"
m["age"] = 18
m["sex"] = true
m["hobby"] = [...]string{
"篮球", "足球", "乒乓球"
}

result, _: = json.Marshal(m)
resultFormatted, _: = json.MarshalIndent(m, "", "	")
fmt.Println("result=", string(result))
fmt.Println("resultFormatted=", string(resultFormatted))
}
/*
result= {"age":18,"hobby":["篮球","足球","乒乓球"],"name":"小王子","sex":true}
resultFormatted= {
        "age": 18,
        "hobby": [
                "篮球",
                "足球",
                "乒乓球"
        ],
        "name": "小王子",
        "sex": true
}
*/
```

实际应用中,我们一般会使用struct字段来定义需要转换的对象,然后再将struct转换为Json格式的数据.

```go
type Person struct {
Name   string   `json:"姓名"`
Age    int      `json:"-"`
Gender bool     `json:"性别"`
Class  int      `json:"班级,omitempty"`
Hobby  []string `json:"爱好"`
}

func main() {
p1 := Person{
Name:   "小明",
Age:    18,
Gender: true,
Class:  1,
Hobby:  []string{"唱", "跳", "rap", "篮球"},
}
result, _ := json.MarshalIndent(p1, "", "  ")
fmt.Println(string(result))
}
```

在定义struct时,可以在字段后添加标签来控制编解码的过程,标签的格式为`key:"value"`,多个标签之间使用空格分隔.

- `json:"字段名"`:用于指定字段对应的Json键名,如果不指定,则默认使用字段名作为Json键名.
- `json:"-"`:用于指定字段不参与编码.
- `json:"字段名,omitempty"`:用于指定字段在编码时忽略空值,即空值不会被编码到Json中,且字段名可以省略.

#### Json解码

json包中的Unmarshal方法用于将json格式数据解码为Go语言中的数据类型,其定义如下:

```go
func Unmarshal(data []byte, v interface{}) error
```

Unmarshal函数将解码后的结果存入v指向的值中.如果v是一个指针,则Unmarshal会将解码后的结果存入v指向的值中,如果v是一个值,则Unmarshal会将解码后的结果存入v的副本中.

函数会将数据解码为如下类型并写入接口:

| Go数据类型                 | Json类型   |
|------------------------|----------|
| bool                   | booleans |
| float64                | numbers  |
| int64                  | numbers  |
| string                 | strings  |
| []interface{}          | arrays   |
| map[string]interface{} | JSON 对象  |
| nil                    | null     |

要将Json数据解码为struct,需要注意以下几点:

- Json数据中的键名必须与struct中的字段名完全匹配,包括大小写.
- Json数据中的键名必须是struct中字段名的子集,否则会报错.

使用map并通过类型断言来进行json解析代码

```go
jsonStr :=
`{
		"name": "小张",
		"age": 30,
		"hobby": ["唱","跳","rap","篮球"],
		"pet":{"petName":"小花","petAge":3}
	}`
// jsonStr是string类型
m := make(map[string]interface{}, 8)
err := json.Unmarshal([]byte(jsonStr), &m)
if err != nil {
fmt.Println("err=", err)
return
}
fmt.Println(m)
for key, value := range m {
switch data := value.(type) {
case string:
fmt.Printf("map[%s]的类型为string,值为%s\n", key, data)
case float64:
fmt.Printf("map[%s]的类型为float64,值为%f\n", key, data)
case []interface{}:
fmt.Printf("map[%s]的类型为[]interface{},值为%v\n", key, data)
case map[string]interface{}:
fmt.Printf("map[%s]的类型为json,值为%v\n", key, data)
case bool:
fmt.Printf("map[%s]的类型为bool,值为%v\n", key, data)
}
}
```

使用struct进行json解码

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string                 `json:"name"`
	Age   int                    `json:"age"`
	Hobby []string               `json:"hobby"`
	Pet   map[string]interface{} `json:"pet"`
}

func (p Person) String() string {
	//return fmt.Sprintf("Person's attribute\nname:%v\nage:%v\nhobby:%v\npet:%v", p.Name, p.Age, p.Hobby, p.Pet)
	return fmt.Sprintf("Person's attribute\nname:%v\nage:%v\nhobby:%v\npet's name:%v\npet's age:%v", p.Name, p.Age, p.Hobby, p.Pet["petName"], p.Pet["petAge"])
}

func main() {
	jsonStr :=
		`{
		"name": "小张",
		"age": 30,
		"hobby": ["唱","跳","rap","篮球"],
		"pet":{"petName":"小花","petAge":3}
	}`
	var p Person
	_ = json.Unmarshal([]byte(jsonStr), &p)
	fmt.Println(p)
}

```

Github开源了一个比标准库解析速度快近10倍的[fastjson库](https://github.com/valyala/fastjson),可以使用这个库来解析json数据.

## 协程与通道

随着社会需求的发展,计算机的性能也在不断提升,单核CPU的性能也在不断提升,但是单核CPU的性能提升的速度已经远远跟不上社会需求的发展,所以我们需要使用多核CPU来提升计算机的性能,但是多核CPU的核心数量是有限的,所以我们需要使用多个CPU来提升计算机的性能,这就是并行.

如今为止,基本上所有的语言都支持并发编程,而Go语言最大的特点就是从语言层面对并发进行了支持.

Go语言主要使用了`协程(coroutine)`和`通道(channel)`来解决并发问题.

### 并发与并行

在计算机程序中,存在很多耗时任务,例如上传文件,下载文件等,这些任务的执行时间往往比较长,如果一个这些任务被一些线程独占,当我们同时需要进行别的任务时,其他任务只能通过等待来执行,需要使用并发来解决这个问题.

并发编程的概念较为广泛,包括多线程编程,多进程编程以及分布式程序等,协程属于多线程编程范畴内.

并发意味着程序可以有多个上下文和多个调用栈.

- `并发(concurrency)`是指同一时间段内,程序可以执行多个任务,把任务在不同的时间点交给处理器进行处理,对于CPU而言,同一时间点并不会同时执行多个任务,而对于人而言,好像是多个任务同时执行. 例如在单核电脑中,可以同时打开qq,浏览器等程序,但是CPU只能同时处理一个任务,所以这些任务是交替执行的,这就是并发.

- `并行(parallelism)`是指把每一个任务分配给每一个处理器独立完成。对于CPU而言,同一时间点可以同时执行多个任务,这就是并行.
