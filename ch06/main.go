package main

import (
	"errors"
	"fmt"
)

// 结构体
type person struct {
	name string
	age  uint16
	//作为字段
	addr address
	//组合
	//类型组合后，外部类型不仅可以使用内部类型的字段，也可以使用内部类型的方法，就像使用自己的方法一样。如果外部类型定义了和内部类型同样的方法，那么外部类型的会覆盖内部类型，这就是方法的覆写
	address
}
type address struct {
	province string
	city     string
}

func (p person) print() {
	fmt.Println(p)
}

// 接口
type Stringer interface {
	String() string
}

//当值类型作为接收者时，person 类型和*person类型都实现了该接口。
//当指针类型作为接收者时，只有*person类型实现了该接口。

// 接口实现
// 方法签名一致 就算实现
func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

//func (p *person) String()  string{
//	return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
//}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

// 工厂函数
func NewPerson(name string) *person {
	return &person{name: name}
}

// 定义接口
type WalkRun interface {
	Walk()
	Run()
}

// 实现接口
func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}

func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}

func main() {

	p := person{name: "bob", age: 58}
	p.print()

	//p2 := person{"bob", 58}
	//fmt.Println(p2.name, p2.age)

	p.address = address{"广东", "深圳"}
	p.addr = p.address

	fmt.Println(p)
	fmt.Println(p.addr)
	fmt.Println(p.city)

	fmt.Println(p.String())
	fmt.Println("===============")
	printString(p)
	printString(&p.address)

	p3 := NewPerson("椰子片")
	fmt.Println(p3)

	err := errors.New("error message")
	fmt.Println(err.Error())

	var s fmt.Stringer
	s = p
	//类型断言
	//判断一个接口的值是否是实现该接口的某个具体类型
	p2, ok := s.(person)
	if ok {
		fmt.Println(p2)
	} else {
		fmt.Println("case error")
	}

}
