package main

import (
	"fmt"
	"unicode/utf8"
)
import "math/rand"

func main() {

	fmt.Println("rand int = ", rand.Int())

	sum := 0

	for i := 1; i < 100; i++ {

		if i%2 != 0 {

			continue

		}

		sum += i

	}

	fmt.Println("the sum is", sum)

	// 数组

	array := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(array[2])

	// 元素初始化全部设置好，可以省略长度
	array = [...]string{"f", "g", "h", "i", "j"}

	fmt.Println(array[2])

	// 初始化特定下标元素
	array1 := [5]string{1: "b", 3: "d"}

	fmt.Println(array1[1])

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}
	fmt.Println("=============for range==================")
	//  for range
	for i, v := range array {

		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)

	}

	// 忽略index
	for _, v := range array {

		fmt.Printf("对应值:%s\n", v)

	}

	// Slice（切片）
	// 基于数组生成切片，包含索引start，但是不包含索引end
	// slice:=array[start:end]

	// 这里是包含索引 2，但是不包含索引 5 的元素，即在 : 右边的数字不会被包含。

	slice := array[2:5]

	fmt.Println(slice)

	// array[:4] 等价于 array[0:4]。
	// array[1:] 等价于 array[1:5]。
	// array[:] 等价于 array[0:5]。

	// 切片声明 切片的容量不能比切片的长度小。
	// 在创建新切片的时候，最好要让新切片的长度和容量一样，
	// 这样在追加操作的时候就会生成新的底层数组，从而和原有数组分离，就不会因为共用底层数组导致修改内容的时候影响多个切片。

	// 声明了一个元素类型为 string 的切片，长度是 4
	// slice1:=make([]string,4)

	// 指定了新创建的切片 []string 容量为 8：
	// slice1:=make([]string,4,8)

	slice1 := []string{"a", "b", "c", "d", "e"}

	fmt.Println(len(slice1), cap(slice1))

	//追加一个元素

	slice2 := append(slice1, "f")

	//多加多个元素

	// slice2:=append(slice1,"f","g")

	//追加另一个切片

	// slice2:=append(slice1,slice...)

	fmt.Println(len(slice2))

	// Map 声明初始化
	nameAgeMap := make(map[string]int)

	// nameAgeMap2:=map[string]int{}
	// nameAgeMap2:=map[string]int{"飞雪无情":20}

	//添加键值对或者更新对应 Key 的 Value
	nameAgeMap["飞雪无情"] = 20
	//获取指定 Key 对应的 Value
	// age:=nameAgeMap["飞雪无情"]

	// Go 语言的 map 可以获取不存在的 K-V 键值对，如果 Key 不存在，返回的 Value 是该类型的零值，比如 int 的零值就是 0。所以很多时候，我们需要先判断 map 中的 Key 是否存在。

	// map 的 [] 操作符可以返回两个值：

	// 第一个值是对应的 Value；
	// 第二个值标记该 Key 是否存在，如果存在，它的值为 true。

	nameAgeMap["飞雪无情1"] = 20

	age, exist := nameAgeMap["飞雪无情"]

	if exist {
		fmt.Println(age)
	}
	// 删除 map key
	delete(nameAgeMap, "飞雪无情")

	// 遍历 Map
	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}

	// 获取所有key
	keys := make([]string, 0, len(nameAgeMap))
	for k := range nameAgeMap {
		keys = append(keys, k)
	}

	fmt.Println("nameAgeMap len ", len(nameAgeMap))

	//string 是一个不可变的字节序列，可以直接转字节切片[]byte
	str := "0123忘"
	bs := []byte(str)
	fmt.Println(bs)

	fmt.Println(str[0])
	//在 UTF8 编码下，一个汉字对应三个字节
	fmt.Println(len(str))
	//汉字算一个长度计算
	fmt.Println(utf8.RuneCountInString(str))

	for i, v := range str {
		fmt.Println(i, v)
	}

	//二维数组
	array2 := [5][2]string{}
	array2[0][0] = "1"
	array2[0][1] = "2"
	array2[1][1] = "3"
	fmt.Println(array2)

	//二维切片

	slice2__ := [2][4]int{}
	fmt.Println("slice2__ = ", slice2__)

	slice2_ := [][]int{{1}, {11, 12}}
	slice2_[0] = append(slice2_[0], 2)

	fmt.Println(slice2_)

}
