package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int `json:"cal_num1"`
	Num2 int `json:"cal_num2"`
}

func calReflect(cal interface{}) {
	val := reflect.ValueOf(cal)
	typ := reflect.TypeOf(cal)

	// val.NumField() 计算reflect.Value的字段的个数
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		// val.Field(i) 计算第i个字段的值
		fmt.Printf("Field value is %v \n", val.Field(i))
		// typ.Field(i).Tag.Get("json") 计算第i个字段的结构体标签
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field tag is %v \n", tagVal)
		}
	}
	// val.NumMethod() 计算方法的数量
	numOfMethod := val.NumMethod()
	fmt.Println("numOfMethod=", numOfMethod)
	// reflect.Value类型的切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf("luyuan"))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0])
}

func (c Cal) GetSub(name string) string {
	str := fmt.Sprintf("%v 完成了减法运算 %v - %v = %v", name, c.Num1, c.Num2, c.Num1-c.Num2)
	return str
}

func main() {
	cal := Cal{10, 1}
	calReflect(cal)
}
