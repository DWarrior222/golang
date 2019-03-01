package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"people_age"`
}

func (this People) Sum(n1, n2 int) int {
	fmt.Println("n1=", n1, "n2=", n2, "n1 + n2 =", n1+n2)
	return n1 + n2
}

func (this People) Set(name string, age int) {
	this.Name = name
	this.Age = age
}

func (this People) Print() {
	fmt.Println("---start~----")
	fmt.Println(this)
	fmt.Println("---end~----")
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	fmt.Println("kd=", kd)
	if kd != reflect.Struct {
		fmt.Println("unexpect struct")
		return
	}

	// 获取结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d field \n", num)

	// 遍历结构体所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: value is %v \n", i, val.Field(i))
		//获取到 struct 标签, 注意需要通过 reflect.Type 来获取 tag 标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag is %v \n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethods := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethods)

	val.Method(0).Call(nil) //获取到第二个方法。调用它

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(2).Call(params)
	fmt.Println("res value is ", res[0].Int()) //返回的结果是 []reflect.Value
}

func main() {
	p := People{
		Name: "luyuan",
		Age:  23,
	}
	TestStruct(p)
}
