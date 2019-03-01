package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的 type , kind, 值
	rType := reflect.TypeOf(b)
	fmt.Println("rType is ", rType)

	rValue := reflect.ValueOf(b)
	fmt.Println("rValue is ", rValue)

	n2 := 2 + rValue.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rValue type is %T, rValue value is %v \n", rValue, rValue)

	iValue := rValue.Interface()
	num2 := iValue.(int)

	fmt.Println("Num2 is", num2)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType is ", rType)

	rValue := reflect.ValueOf(b)
	fmt.Println("rValue is ", rValue)

	iValue := rValue.Interface()
	fmt.Printf("iValue type is %T, iValue value is %v \n", iValue, iValue)

	stu, ok := iValue.(Student)

	if ok {
		fmt.Println("stu.Name is ", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val type is %T \n", val)
	val.Elem().SetInt(100)
	fmt.Printf("val is %v \n", val)
}

func main() {
	n := 100
	reflectTest01(n)

	stu := Student{
		Name: "luyuan",
		Age:  23,
	}
	reflectTest02(stu)

	n1 := 20
	testInt(&n1)
	fmt.Println(n1)
}
