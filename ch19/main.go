package main

import (
	"fmt"
	"os"
)

//TODO 性能优化：Go 语言如何进行代码检查和优化？

// 保证代码质量和性能的手段不只有单元测试和基准测试，还有代码规范检查和性能优化。
// 代码规范检查是对单元测试的一种补充，它可以从非业务的层面检查你的代码是否还有优化的空间，比如变量是否被使用、是否是死代码等等。
// 性能优化是通过基准测试来衡量的，这样我们才知道优化部分是否真的提升了程序的性能。

//golangci-lint
//golangci-lint 是一个集成工具，它集成了很多静态代码分析工具，便于我们使用。通过配置这一工具，我们可以很灵活地启用需要的代码规范检查。
//可用于 Go 语言代码分析的工具有很多，比如 golint、gofmt、misspell 等，如果一一引用配置，就会比较烦琐，所以通常我们不会单独地使用它们，而是使用 golangci-lint。

//go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
//golangci-lint version

//golangci-lint run ch19/

//golangci-lint 配置
// 默认启用的 linter
//deadcode - 死代码检查
//errcheck - 返回错误是否使用检查
//gosimple - 检查代码是否可以简化
//govet - 代码可疑检查，比如格式化字符串和类型不一致
//ineffassign - 检查是否有未使用的代码
//staticcheck - 静态分析检查
//structcheck - 查找未使用的结构体字段
//typecheck - 类型检查
//unused - 未使用代码检查
//varcheck - 未使用的全局变量和常量检查

//查看 linter 列表
//golangci-lint linters

//修改默认启用的 linter，就需要对 golangci-lint 进行配置。即在项目根目录下新建一个名字为 .golangci.yml 的文件

//Go 语言有两部分内存空间：栈内存和堆内存。
//
//栈内存由编译器自动分配和释放，开发者无法控制。栈内存一般存储函数中的局部变量、参数等，函数创建的时候，这些内存会被自动创建；函数返回的时候，这些内存会被自动释放。
//堆内存的生命周期比栈内存要长，如果函数返回的值还会在其他地方使用，那么这个值就会被编译器自动分配到堆上。堆内存相比栈内存来说，不能自动被编译器释放，只能通过垃圾回收器才能释放，所以栈内存效率会很高。

//逃逸分析
//逃逸分析是判断变量是分配在堆上还是栈上的一种方法，在实际的项目中要尽可能避免逃逸，这样就不会被 GC 拖慢速度，从而提升效率。
// 从逃逸分析来看，指针虽然可以减少内存的拷贝，但它同样会引起逃逸，所以要根据实际情况选择是否使用指针。

//go build -gcflags="-m -l" ./ch19/server_main.go
//-m 表示打印出逃逸分析信息，-l 表示禁止内联，可以更好地观察逃逸。从以上输出结果可以看到，发生了逃逸，也就是说指针作为函数返回值的时候，一定会发生逃逸。
//逃逸到堆内存的变量不能马上被回收，只能通过垃圾回收标记清除，增加了垃圾回收的压力，所以要尽可能地避免逃逸，让变量分配在栈内存上，这样函数返回时就可以回收资源，提升效率。

//优化技巧  性能优化的时候，要结合基准测试，来验证自己的优化是否有提升。
//第 1 个需要介绍的技巧是尽可能避免逃逸，因为栈内存效率更高，还不用 GC。比如小对象的传参，array 要比 slice 效果好。
//如果避免不了逃逸，还是在堆上分配了内存，那么对于频繁的内存申请操作，我们要学会重用内存，比如使用 sync.Pool，这是第 2 个技巧。
//第 3 个技巧就是选用合适的算法，达到高性能的目的，比如空间换时间。

//尽可能避免使用锁、并发加锁的范围要尽可能小、使用 StringBuilder 做 string 和 [ ] byte 之间的转换、defer 嵌套不要太多

//TODO 性能剖析工具 pprof

const name = "飞雪无情"

func main() {

	os.Mkdir("tmp", 0666)

	//字符串逃逸到了堆上，这是因为「hello」这个字符串被已经逃逸的指针变量引用，所以它也跟着逃逸了
	//被已经逃逸的指针引用的变量也会发生逃逸

	fmt.Println("hello")

	//变量 m 没有逃逸，反而被变量 m 引用的变量 s 逃逸到了堆上。所以被map、slice 和 chan 这三种类型引用的指针一定会发生逃逸的。
	m := map[int]*string{}
	s := "飞雪无情"
	m[0] = &s
}

// 指针作为函数返回逃逸的例子
func newString() *string {

	s := new(string)

	*s = "飞雪无情"

	return s

}

// 优化后
func newString2() string {

	s := new(string)

	*s = "飞雪无情"

	return *s

}
