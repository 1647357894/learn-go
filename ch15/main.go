package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// 运行时反射

//反射修改一个值的规则
//可被寻址，通俗地讲就是要向 reflect.ValueOf 函数传递一个指针作为参数。
//如果要修改 struct 结构体字段值的话，该字段需要是可导出的，而不是私有的，也就是该字段的首字母为大写。
//记得使用 Elem 方法获得指针指向的值，这样才能调用 Set 系列方法进行修改。

//反射的三大定律
//任何接口值 interface{} 都可以反射出反射对象，也就是 reflect.Value 和 reflect.Type，通过函数 reflect.ValueOf 和 reflect.TypeOf 获得。
//反射对象也可以还原为 interface{} 变量，也就是第 1 条定律的可逆性，通过 reflect.Value 结构体的 Interface 方法获得。
//要修改反射的对象，该值必须可设置，也就是可寻址，参考上节课修改变量的值那一节的内容理解。
//小提示：任何类型的变量都可以转换为空接口 intferface{}，所以第 1 条定律中函数 reflect.ValueOf 和 reflect.TypeOf 的参数就是 interface{}，表示可以把任何类型的变量转换为反射对象。在第 2 条定律中，reflect.Value 结构体的 Interface 方法返回的值也是 interface{}，表示可以把反射对象还原为对应的类型变量。

func main() {
	i := 3

	iv := reflect.ValueOf(i)

	it := reflect.TypeOf(i)

	fmt.Println(iv, it) //3 int

	//reflect.Value 和 int 类型互转
	i1 := iv.Interface().(int)
	fmt.Println(i1)

	//反射修改值
	ipv := reflect.ValueOf(&i)
	ipv.Elem().SetInt(4)
	fmt.Println(i)
	p := person{Name: "无情", Age: 20}

	ppv := reflect.ValueOf(&p)

	elem := ppv.Elem()
	elem.Field(0).SetString("张三")
	fmt.Println(p)

	fmt.Println(ppv.Kind())
	pv := reflect.ValueOf(p)
	//底层类型
	fmt.Println(pv.Kind())

	pt := reflect.TypeOf(p)
	//遍历person的字段
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}

	//遍历person的方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}

	//反射调用方法
	stringMethod := pv.MethodByName("String")
	//args:=[]reflect.Value{reflect.ValueOf("登录")}
	fmt.Println("call method")
	callResult := stringMethod.Call([]reflect.Value{})
	fmt.Println("call method result ", callResult[0])

	//反射调用函数

	//判断是否实现某接口
	//尽可能通过类型断言的方式判断是否实现了某接口，而不是通过反射。
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：", pt.Implements(stringerType))
	fmt.Println("是否实现了io.Writer：", pt.Implements(writerType))

	//	json convert
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}

	//json to struct
	//respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	respJSON := "{\"name\":\"李四\",\"age\":40}"
	err = json.Unmarshal([]byte(respJSON), &p)
	if err == nil {
		fmt.Println(p)
	}

	//遍历person字段中key为json的tag

	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json"))
	}

	struct2Json(p)

}

type person struct {
	//Struct Tag
	//struct tag 是一个添加在 struct 字段上的标记，使用它进行辅助，可以完成一些额外的操作，比如 json 和 struct 互转
	//结构体的字段可以有多个 tag，用于不同的场景，比如 json 转换、yaml 解析、orm 解析等。如果有多个 tag，要使用空格分隔
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
}

func struct2Json(p person) {
	pv := reflect.ValueOf(p)
	pt := reflect.TypeOf(p)
	//自己实现的struct to json
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	num := pt.NumField()
	for i := 0; i < num; i++ {
		jsonTag := pt.Field(i).Tag.Get("json") //获取json tag
		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		//获取字段的值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))
		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}

	jsonBuilder.WriteString("}")
	fmt.Println("---------------------")
	fmt.Println(jsonBuilder.String()) //打印json字符串
}
