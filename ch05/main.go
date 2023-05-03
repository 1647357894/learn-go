package main

import (
	"errors"
	"fmt"
)

// 函数和方法
// 函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用；
// 函数名称首字母大写代表公有函数，不同的包也可以调用；
// 任何一个函数都会从属于一个包。

// 在 Go 语言中，虽然存在函数和方法两个概念，但是它们基本相同，不同的是所属的对象。函数属于一个包，方法属于一个类型，所以方法也可以简单地理解为和一个类型关联的函数。
func main() {
	fmt.Println("hello")

	result := sum(1, 2)
	fmt.Println(result)
	result3, err := sum3(0, 1)
	if err == nil {
		fmt.Println(result3)
	} else {
		fmt.Println(err)
	}

	fmt.Println(sumAll(1, 2, 3))

	//匿名函数
	sum5 := func(x, y int) int { return x + y }
	fmt.Println(sum5(1, 2))

	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	//	方法调用
	age := Age(25)
	age.String("lucky")
	age.modify()
	age.String("lk")

	//方法赋值给变量，方法表达式
	sm := Age.String
	//通过变量，要传一个接收者进行调用也就是age
	sm(age, "hello")
}

// 定义函数
func sum(x int, y int) int {
	return x + y
}

// 参数类型一样，可以省略一个类型定义
func sum2(x, y int) int {
	return x + y
}

// 多值返回
func sum3(x, y int) (int, error) {
	if x <= 0 {
		return 0, errors.New("x less or equal 0")
	}
	return x + y, nil
}

// 返回参数命名
func sum4(x, y int) (result int, err error) {
	if x <= 0 {
		return 0, errors.New("x less or equal 0")
	}
	result = x + y
	err = nil
	//直接为命名参数赋值，可以省略return返回值
	return
}

func sum4_(x, y int) (int, error) {
	if x <= 0 {
		return 0, errors.New("x less or equal 0")
	}
	result := x + y
	return result, nil
}

func voidMethod() {
	fmt.Println("void method")
}

func Println(a ...any) (n int, err error) {
	return fmt.Println(a)
}

func sumAll(params ...int) (sum int) {
	for _, param := range params {
		sum += param
	}
	return
}

// 闭包函数
func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 方法
// 定义类型
type Age uint

// 方法绑定 Age类型
func (age Age) String(name string) {
	fmt.Println("name is ", name)
	fmt.Println("age is ", age)
}
func (age *Age) modify() {
	*age = Age(32)
}
func (age Age) modify2() {
	age = Age(66)
}
