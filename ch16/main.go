package main

import (
	"fmt"
	"unsafe"
)

// unsafe

// unsafe.Pointer
// unsafe.Pointer 是一种特殊意义的指针，可以表示任意类型的地址，类似 C 语言里的 void* 指针，是全能型的。

//uintptr 指针类型
//uintptr 也是一种指针类型，它足够大，可以表示任何指针
//因为 unsafe.Pointer 不能进行运算，比如不支持 +（加号）运算符操作，但是 uintptr 可以。通过它，可以对指针偏移进行计算，这样就可以访问特定的内存，达到对特定内存读写的目的，这是真正内存级别的操作。

// Go 语言中存在三种类型的指针，它们分别是：常用的 *T、unsafe.Pointer 及 uintptr
//任何类型的 *T 都可以转换为 unsafe.Pointer；
//unsafe.Pointer 也可以转换为任何类型的 *T；
//unsafe.Pointer 可以转换为 uintptr；
//uintptr 也可以转换为 unsafe.Pointer。

//unsafe.Pointer 主要用于指针类型的转换，而且是各个指针类型转换的桥梁。uintptr 主要用于指针运算，尤其是通过偏移量定位不同的内存。

// unsafe.Sizeof
// Sizeof 函数可以返回一个类型所占用的内存大小，这个大小只与类型有关，和类型对应的变量存储的内容大小无关，比如 bool 型占用一个字节、int8 也占用一个字节。
// 一个 struct 结构体的内存占用大小，等于它包含的字段类型内存占用大小之和。
type person struct {
	Name string
	Age  int
}

func main() {

	//convertMethod1()

	//类型转换
	// *int 转换为 *float64
	//convertMethod2()

	//先使用 new 函数声明一个 *person 类型的指针变量 p。
	//然后把 *person 类型的指针变量 p 通过 unsafe.Pointer，转换为 *string 类型的指针变量 pName。
	//因为 person 这个结构体的第一个字段就是 string 类型的 Name，所以 pName 这个指针就指向 Name 字段（偏移为 0），对 pName 进行修改其实就是修改字段 Name 的值。
	//因为 Age 字段不是 person 的第一个字段，要修改它必须要进行指针偏移运算。所以需要先把指针变量 p 通过 unsafe.Pointer 转换为 uintptr，这样才能进行地址运算。既然要进行指针偏移，那么要偏移多少呢？这个偏移量可以通过函数 unsafe.Offsetof 计算出来，该函数返回的是一个 uintptr 类型的偏移量，有了这个偏移量就可以通过 + 号运算符获得正确的 Age 字段的内存地址了，也就是通过 unsafe.Pointer 转换后的 *int 类型的指针变量 pAge。
	//然后需要注意的是，如果要进行指针运算，要先通过 unsafe.Pointer 转换为 uintptr 类型的指针。指针运算完毕后，还要通过 unsafe.Pointer 转换为真实的指针类型（比如示例中的 *int 类型），这样可以对这块内存进行赋值或取值操作。
	//有了指向字段 Age 的指针变量 pAge，就可以对其进行赋值操作，修改字段 Age 的值了。

	p := new(person)
	//Name是person的第一个字段不用偏移，即可通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "飞雪无情"
	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)

	fmt.Println("========================")
	fmt.Println(unsafe.Sizeof(true))

	fmt.Println(unsafe.Sizeof(int8(0)))

	fmt.Println(unsafe.Sizeof(int16(10)))

	fmt.Println(unsafe.Sizeof(int32(10000000)))

	fmt.Println(unsafe.Sizeof(int64(10000000000000)))

	fmt.Println(unsafe.Sizeof(int(10000000000000000)))

	fmt.Println(unsafe.Sizeof(string("飞雪无情")))

	fmt.Println(unsafe.Sizeof([]string{"飞雪u无情", "张三"}))

	//unsafe.Alignof()

}

func convertMethod2() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
}

func convertMethod1() {
	//i := 10
	//
	//ip := &i
	//
	//var fp *float64 = (*float64)(ip)
	//
	//fmt.Println(fp)

}
