package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 通过 error、deferred、panic 等处理错误

type commonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息
}

// implement error interface
// *commonError指针类型
func (ce *commonError) Error() string {
	return ce.errorMsg
}

func add(x, y int) (int, error) {
	if x < 0 || y < 0 {
		//&commonError取该对象的指针,指针变量
		return 0, &commonError{
			errorCode: 1,
			errorMsg:  "a或者b不能为负数"}
	}
	return x + y, nil
}

type Person struct {
	name string
	age  int
}

func (r *Person) hello() string {
	return r.name
}

func hello(name string) {
	if name == "" {
		panic("name不能为空")
	}
	fmt.Println("hello ", name)
	//省略其他代码
}

// 在一个方法或者函数中，可以有多个 defer 语句；
// 多个 defer 语句的执行顺序依照后进先出的原则。
func moreDefer() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Three defer")
	fmt.Println("函数自身代码")
}

func main() {

	//person 表示指针变量
	//*person 表示指针对象存放的内容
	//&person 表示取指针变量的地址

	person := &Person{"liu", 10} //&是取地址符号, 取到Person类型对象的地址
	fmt.Println(person)          //&{liu 10} *可以表示一个变量是指针类型(person是一个指针变量)
	fmt.Println(*person)         //{liu 10}  *也可以表示指针类型变量所指向的存储单元 ,也就是这个地址所指向的值
	fmt.Println(&person)         //0xc000006030  查看这个指针变量的地址 , 基本数据类型直接打印地址
	var person2 *Person = &Person{"liu2", 10}
	fmt.Println(person2) //&{liu2 10}
	fmt.Println("------------------------")

	r, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r2, err2 := add(-1, 2)
	if err2 == nil {
		fmt.Println(r2)
	}

	fmt.Println(err2)

	//类型断言
	//* 表示取指针地址中的值
	//& 表示取一个值的地址

	//& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
	//*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
	if cm, ok := err2.(*commonError); ok {
		fmt.Println("errcode=", cm.errorCode, "errorMsg=", cm.errorMsg)
		fmt.Println(cm.Error())
	}

	//warp error
	oneErr := errors.New("one error")
	warpErr := fmt.Errorf("warp  %w", oneErr)
	fmt.Println(oneErr)
	fmt.Println(warpErr)
	//Unwrap error
	fmt.Println(errors.Unwrap(warpErr))

	//比较error是否相等
	//两个 error 相等或 err 包含 target 的情况下返回 true，其余返回 false
	fmt.Println(errors.Is(warpErr, oneErr))

	//error 断言

	var cm *commonError
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	}

	if errors.As(warpErr, &oneErr) {

	}
	//	在 Go 语言提供的 Error Wrapping 能力下，我们写的代码要尽可能地使用 Is、As 这些函数做判断和转换

	//Deferred 函数

	//defer关键字保证方法执行结束后，该关键字修饰的代码一定会执行
	//defer 语句常被用于成对的操作，如文件的打开和关闭，加锁和释放锁，连接的建立和断开等。不管多么复杂的操作，都可以保证资源被正确地释放。

	defer fmt.Println("defer exec")

	file, err := os.ReadFile("C://tmp//test.txt")
	if err == nil {
		fmt.Println("file hex ", hex.Dump(file))
	}

	//Recover 捕获 Panic 异常
	//通常情况下，我们不对 panic 异常做任何处理，因为既然它是影响程序运行的异常，就让它直接崩溃即可。但是也的确有一些特例，比如在程序崩溃前做一些资源释放的处理，这时候就需要从 panic 异常中恢复，才能完成处理。
	//通过内置的 recover 函数恢复 panic 异常。因为在程序 panic 异常崩溃的时候，只有被 defer 修饰的函数才能被执行，所以 recover 函数要结合 defer 关键字使用才能生效。
	//recover 函数返回的值就是通过 panic 函数传递的参数值。
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	moreDefer()

	//	Panic 异常
	//Go 语言是一门静态的强类型语言，很多问题都尽可能地在编译时捕获，但是有一些只能在运行时检查，比如数组越界访问、不相同的类型强制转换等，这类运行时的问题会引起 panic 异常。
	//除了运行时可以产生 panic 外，我们自己也可以抛出 panic 异常
	//panic 异常是一种非常严重的情况，会让程序中断运行，使程序崩溃，所以如果是不影响程序运行的错误，不要使用 panic，使用普通错误 error 即可。
	hello("")
	//panic异常后面的代码使用 defer也不会执行
	fmt.Println("defer exec again")

}
