# 基础语法部分

## 关键字

Go语言中共有`25`个关键字，如下所示：

```
break       default     func    interface   select
case        defer       go      map         struct 
chan        else        goto    package     switch
const       fallthrough if      range       type
continue    for         import  return      var
```

此外,包含`36`个预定义标识符,如下所示:

```
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

### 变量声明

Go语言中,变量需要声明后才可以使用,同一作用域内不支持重复声明,且声明后必须使用,否则会报错.  
Go语言推荐驼峰式命名法,例如: `userName`  
声明代码如下:

```
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
```
