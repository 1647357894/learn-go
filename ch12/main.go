package main

import "fmt"

// 指针
// 在编程语言中，指针是一种数据类型，用来存储一个内存地址，该地址指向存储在该内存中的对象。这个对象可以是字符串、整数、函数或者你自定义的结构体。
// 在 Go 语言中使用类型名称前加 * 的方式，即可表示一个对应的指针类型。比如 int 类型的指针类型是 *int，float64 类型的指针类型是 *float64，自定义结构体 A 的指针类型是 *A。总之，指针类型就是在对应的类型前加 * 号

//指针接受者
//如果接收者类型是 map、slice、channel 这类引用类型，不使用指针；
//如果需要修改接收者，那么需要使用指针；
//如果接收者是比较大的类型，可以考虑使用指针，因为内存拷贝廉价，所以效率高。

//用指针的好处
//可以修改指向数据的值；
//在变量赋值，参数传值的时候可以节省内存。

//不要对 map、slice、channel 这类引用类型使用指针；
//如果需要修改方法接收者内部的数据或者状态时，需要使用指针；
//如果需要修改参数的值或者内部数据时，也需要使用指针类型的参数；
//如果是比较大的结构体，每次参数传递或者调用方法都要内存拷贝，内存占用多，这时候可以考虑使用指针；
//像 int、bool 这样的小数据类型没必要使用指针；
//如果需要并发安全，则尽可能地不要使用指针，使用指针一定要保证并发安全；
//指针最好不要嵌套，也就是不要使用一个指向指针的指针，虽然 Go 语言允许这么做，但是这会使你的代码变得异常复杂。

func main() {

	name := "jack"
	age := 18

	//指针变量的类型为 *string
	//指针类型变量 只占用 4 个或者 8 个字节的内存大小
	nameP := &name //取地址

	var intP *int
	intP = &age //指针类型不同，无法赋值

	fmt.Println("name变量的值为:", name)

	fmt.Println("name变量的内存地址为:", nameP)
	//通过指针变量获取具体值
	nameV := *nameP
	fmt.Println("nameP指针指向的值为:", nameV)

	//修改指针指向的值
	*nameP = "yey"
	fmt.Println("update nameP指针指向的值为:", *nameP)
	fmt.Println("name值为:", name)

	fmt.Println(age)
	fmt.Println(intP)
	fmt.Println(*intP)

	intP1 := new(int)
	fmt.Println(intP1)

	//通过var关键字定义的指针变量不能直接赋值，因为它的值还是nil，也就是还没有指向内存地址，需要使用new函数分配内存
	var intP2 *int = new(int)

	*intP2 = 10

	fmt.Println(intP2)

	modifyAge(&age)

	fmt.Println("age的值为:", age)

}

// age 只是实参 age 的一份拷贝，所以修改它不会改变实参 age 的值
//	func modifyAge(age int) {
//		age = 20
//	}

// 在函数中通过形参改变实参的值时，需要使用指针类型的参数
func modifyAge(age *int) {
	*age = 20
}
