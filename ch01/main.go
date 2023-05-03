package main

import (
	"fmt"
	"strconv"
	"strings"

	"rsc.io/quote"
)

func main() {

	// 有符号整型：如 int、int8、int16、int32 和 int64。
	// 无符号整型：如 uint、uint8、uint16、uint32 和 uint64。
	// byte

	// 两种精度的浮点数，分别是 float32 和 float64

	// 未使用的变量 编译时会报错

	var (
		j = 0
		k = 1
	)
	fmt.Println("Hello, 世界")
	var i int = 10

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	var f32 float32 = 2.2

	var f64 float64 = 10.3456

	fmt.Println("f32 is", f32, ",f64 is", f64)

	var bf bool = false

	var bt = true

	fmt.Println("bf is", bf, ",bt is", bt)

	var s1 string = "Hello"

	var s2 string = "世界"

	fmt.Println("s1 is", s1, ",s2 is", s2)

	fmt.Println("s1+s2=", s1+s2)

	var zi int

	var zf float64

	var zb bool

	var zs string

	fmt.Println(zi, zf, zb, zs)

	// 变量简短声明   变量名:=表达式
	s11 := "Hello"
	fmt.Println(s11)

	// 在 Go 语言中，指针对应的是变量在内存中的存储位置，也就说指针的值就是变量的内存地址。通过 & 可以获取一个变量的地址，也就是指针。
	// 在以下的代码中，pi 就是指向变量 i 的指针。要想获得指针 pi 指向的变量值，通过*pi这个表达式即可。尝试运行这段程序，会看到输出结果和变量 i 的值一样。
	pi := &s1
	fmt.Println(*pi)

	// 赋值
	s11 = "3.14"
	fmt.Println(s11)

	// 常量 只允许布尔型、字符串、数字类型这些基础类型作为常量。
	const name = "hello go"

	// iota 是一个常量生成器，它可以用来初始化相似规则的常量
	const (
		one = iota + 1

		two

		three

		four
	)

	fmt.Println(one, two, three, four)

	// 字符串和数字互转

	number1 := 1

	// 数字转字符串
	i2s := strconv.Itoa(number1)
	// 字符串转数字  err 转换异常
	s2i, err := strconv.Atoi(i2s)

	// 字符串转数字  err 转换异常
	s3i, err := strconv.Atoi(s11)

	fmt.Println(i2s, s2i, err, s3i, err)

	fmt.Println(quote.Go())

	//判断s1的前缀是否是H

	fmt.Println(strings.HasPrefix(s1, "H"))

	//在s1中查找字符串o

	fmt.Println(strings.Index(s1, "o"))

	//把s1全部转为大写

	fmt.Println(strings.ToUpper(s1))

}
