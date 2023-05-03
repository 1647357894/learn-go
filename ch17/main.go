package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// slice如何高效处理数据
// 切片是对数组的抽象和封装，它的底层是一个数组存储所有的元素，但是它可以动态地添加元素，容量不足时还可以自动扩容，你完全可以把切片理解为动态数组

//高效的原因
//如果从集合类型的角度考虑，数组、切片和 map 都是集合类型，因为它们都可以存放元素，但是数组和切片的取值和赋值操作要更高效，因为它们是连续的内存操作，通过索引就可以快速地找到元素存储的地址。
//进一步对比，在数组和切片中，切片又是高效的，因为它在赋值、函数传参的时候，并不会把所有的元素都复制一遍，而只是复制 SliceHeader 的三个字段就可以了，共用的还是同一个底层数组。

//切片的高效还体现在 for range 循环中，因为循环得到的临时变量也是个值拷贝，所以在遍历大的数组时，切片的效率更高。
//切片基于指针的封装是它效率高的根本原因，因为可以减少内存的占用，以及减少内存复制时的时间消耗。

func main() {
	///数组类型 由长度和元素类型组成

	//数组的局限性
	//一旦一个数组被声明，它的大小和内部元素的类型就不能改变
	//a1 := [1]string{"飞雪无情"}
	//a2 := [2]string{"飞雪无情"}
	//fmt.Println(a1, a2)

	//如果切片的容量不够，append 函数会自动扩容
	//append 自动扩容的原理是新创建一个底层数组，把原来切片内的元素拷贝到新数组中，然后再返回一个指向新数组的切片。
	ss := []string{"飞雪无情", "张三"}
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))
	ss = append(ss, "李四", "王五")
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))
	fmt.Println(ss)

	a1 := [2]string{"飞雪无情", "张三"}
	s1 := a1[0:1]
	s2 := a1[:]

	//打印出s1和s2的Data值，是一样的
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	sh1 := (*mySlice)(unsafe.Pointer(&s1))
	fmt.Println(sh1.Data, sh1.Len, sh1.Cap)

	//同一个数组在 main 函数中的指针和在 arrayF 函数中的指针是不一样的，这说明数组在传参的时候被复制了，又产生了一个新数组。
	//而 slice 切片的底层 Data 是一样的，这说明不管是在 main 函数还是 sliceF 函数中，这两个切片共用的还是同一个底层数组，底层数组并没有被复制。
	fmt.Printf("函数main数组指针：%p\n", &a1)
	arrayF(a1)
	s1 = a1[0:1]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)

	//	string 和 []byte 互转
	//Go 语言通过先分配一个内存再复制内容的方式，实现 string 和 []byte 之间的强制转换
	//string 和 []byte 类型互转的具体实现 runtime.stringtoslicebyte 和 runtime.slicebytetostring
	s := "飞雪无情"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)

	//s4 没有申请新内存（零拷贝），它和变量 b 使用的是同一块内存，因为它们的底层 Data 字段值相同，这样就节约了内存，也达到了 []byte 转 string 的目的
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Printf("s4的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)).Data)

	//	通过 unsafe.Pointer 把 string 转为 []byte 后，不能对 []byte 修改，比如不可以进行 b1[0]=12 这种操作，会报异常，导致程序崩溃。这是因为在 Go 语言中 string 内存是只读的。
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh))
	fmt.Println(b1)

}
func arrayF(a [2]string) {

	fmt.Printf("函数arrayF数组指针：%p\n", &a)

}

func sliceF(s []string) {

	fmt.Printf("函数sliceF Data：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)

}

type mySlice struct {
	Data uintptr

	Len int

	Cap int
}
