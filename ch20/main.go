package main

//导包
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iswang.jie.com/m/v2/ch20/util"
	"time"
)

// 一个包就是一个独立的空间，你可以在这个包里定义函数、结构体等。这时，我们认为这些函数、结构体是属于这个包的。

//作用域
//Go 语言中，所有的定义，比如函数、变量、结构体等，如果首字母是大写，那么就可以被其他包使用；
//反之，如果首字母是小写的，就只能在同一个包内使用。

// init 函数
// ，Go 语言还有一个特殊的函数——init，通过它可以实现包级别的一些初始化操作。
// init 函数没有返回值，也没有参数，它先于 main 函数执行
// 一个包中可以有多个 init 函数，但是它们的执行顺序并不确定，所以如果你定义了多个 init 函数的话，要确保它们是相互独立的，一定不要有顺序上的依赖
func init() {
	fmt.Println("init in server_main.go ")
}

//Go 语言中的模块
//一个模块通常是一个项目，比如这个专栏实例中使用的 gotour 项目；
//也可以是一个框架，比如常用的 Web 框架 gin。

//go mod

//创建一个模块
//模块名最好是以自己的域名开头，比如 flysnow.org/gotour，这样就可以很大程度上保证模块名的唯一，不至于和其他模块重名。
//go mod init gotour

//引用第三方模块
//设置代理
//go env -w GO111MODULE=on
//go env -w GOPROXY=https://goproxy.io,direct

//# 设置不走 proxy 的私有仓库，多个用逗号相隔（可选）
//go env -w GOPRIVATE=*.corp.example.com

//安装gin
//go get -u github.com/gin-gonic/gin
//同步模块的依赖
//go mod tidy

func main() {

	fmt.Println(time.Now())
	util.Print("success")
	engine := gin.Default()
	engine.Run()
}
