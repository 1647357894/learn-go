package main

import "fmt"

//值类型
//Go 语言中的函数传参都是值传递。 值传递指的是传递原来数据的一份拷贝，而不是原来的数据本身。
//除了 struct 外，还有浮点型、整型、字符串、布尔、数组，这些都是值类型。

//指针类型
//指针类型的变量保存的值就是数据对应的内存地址，所以在函数参数传递是传值的原则下，拷贝的值也是内存地址
//指针类型的参数是永远可以修改原数据的，因为在参数传递时，传递的是内存地址。
//值传递的是指针，也是内存地址。通过内存地址可以找到原数据的那块内存，所以修改它也就等于修改了原数据。

//引用类型
//严格来说，Go 语言没有引用类型，但是我们可以把 map、chan 称为引用类型，这样便于理解。除了 map、chan 之外，Go 语言中的函数、接口、slice 切片都可以称为引用类型。
// 指针类型也可以理解为是一种引用类型。

//在 Go 语言中，函数的参数传递只有值传递，而且传递的实参都是原始数据的一份拷贝。如果拷贝的内容是值类型的，那么在函数中就无法修改原始数据；如果拷贝的内容是指针（或者可以理解为引用类型 map、chan 等），那么就可以在函数中修改原始数据。

type address struct {
	province string
	city     string
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

// 当值类型作为接收者实现了某接口时，它的指针类型也同样实现了该接口
// 虽然指向具体类型的指针可以实现一个接口，但是指向接口的指针永远不可能实现该接口
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

//func modifyPerson(p person)  {
//	p.name = "李四"
//	p.age = 20
//}

func modifyPerson(p *person) {
	p.name = "李四"
	p.age = 20
}

type person struct {
	name string
	age  int
}

func modifyMap(p map[string]int) {
	fmt.Printf("modifyMap函数：p的内存地址为%p\n", p)
	p["飞雪无情"] = 20

}

func main() {
	add := address{province: "北京", city: "北京"}
	printString(add)
	printString(&add)

	var si fmt.Stringer = address{province: "上海", city: "上海"}
	printString(si)
	//sip := &si
	//printString(sip)

	//修改参数
	p := person{name: "张三", age: 18}
	modifyPerson(&p)
	fmt.Println("person name:", p.name, ",age:", p.age)

	//在 Go 语言中，任何创建 map 的代码（不管是字面量还是 make 函数）最终调用的都是 runtime.makemap 函数
	//makemap 函数返回的是一个 *hmap 类型，也就是说返回的是一个指针，所以我们创建的 map 其实就是一个 *hmap。
	m := make(map[string]int)
	fmt.Printf("main函数：m的内存地址为%p\n", m)

	m["飞雪无情"] = 18

	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])

	modifyMap(m)

	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])

}
