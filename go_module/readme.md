# 什么是go module
[教程文章链接](https://www.liwenzhou.com/posts/Go/dependency/)  
`go module`是1.11版本推出的管理工具,从1.13开始,成为默认的依赖管理工具.
## GO111MODULE
要启用`go module`支持首先要设置环境变量`GO111MODULE`，通过它可以开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是auto。
* GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
* GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
* GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。

需要通过命令行进行修改GO111MODULE的值,
```terminal
go env -w GO111MODULE=auto
```
简单来说，设置GO111MODULE=on之后就可以使用`go module`了，以后就没有必要在GOPATH中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

使用 go module 管理依赖后会在项目根目录下生成两个文件go.mod和go.sum。
## go env命令
`go env`命令可选三个flag
* The -json flag prints the environment in JSON format
instead of as a shell script.(以json格式输出)
* The -u flag requires one or more arguments and unsets
the default setting for the named environment variables,
if one has been set with 'go env -w'.(重置某个参数)
* The -w flag requires one or more arguments of the
form NAME=VALUE and changes the default settings
of the named environment variables to the given values.(修改某个参数)
## go mod命令
Go mod provides access to operations on modules.  
Go mod 提供处理依赖的操作方法,包括下述命令
| 命令     | 英文描述                                   | 中文描述                           |
| -------- | ------------------------------------------ | ---------------------------------- |
| download | download modules to local cache            | 将依赖下载到本地缓存               |
| edit     | edit go.mod from tools or scripts          | 编辑go.mod文件                     |
| graph    | print module requirement graph             | 打印模块依赖图                     |
| init     | initialize new module in current directory | 初始化当前文件夹, 创建go.mod文件   |
| tidy     | add missing and remove unused modules      | 增加缺少的module，删除无用的module |
| vendor   | make vendored copy of dependencies         | 将依赖复制到vendor下               |
| verify   | verify dependencies have expected content  | 校验依赖                           |
| why      | explain why packages or modules are needed | 解释为什么需要依赖                 |

## go.mod文件
go.mod文件记录了项目所有的依赖信息,结构大致如下:
```go
module gin_demo

go 1.19

require github.com/gin-gonic/gin v1.8.2

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
    ...
)
```
其中,
* `module`用来定义包名
* `go`用来定义go版本
* `require`用来定义依赖包以及版本
* `indirect`表示间接引用

## go get命令
在项目中,使用`go get`命令可以下载依赖包

* go get example.com/pkg 下载最新版本依赖
* go get example.com/pkg@v1.2.3 下载指定版本的依赖
* go get example.com/mod@none 删除依赖
* 如果下载所有依赖可以使用go mod download命令。

# 在项目中使用go module
## 对现存项目
如果需要对一个已经存在的项目启用`go module`，可以按照以下步骤操作：
1. 在项目目录下执行go mod init，生成一个go.mod文件。
2. 执行go get，查找并记录当前项目的依赖，同时生成一个go.sum记录每个依赖库的版本和哈希值。
## 对新项目
对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：
1. 执行go mod init 项目名命令，在当前项目文件夹下创建一个go.mod文件。
2. 手动编辑go.mod中的require依赖项或执行go get自动发现、维护依赖。

# [本地包调用](./main.go)