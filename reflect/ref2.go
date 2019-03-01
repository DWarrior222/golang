package main

import (
	"fmt"
	"reflect"
)

func testFloat(b interface{}) {
	rValue := reflect.ValueOf(b)
	// rType := reflect.TypeOf(b)
	rType := rValue.Type()
	rKind := rValue.Kind()
	fmt.Printf("rValue value is %v, type is %v, kind is %v\n", rValue, rType, rKind)

	iValue := rValue.Interface()
	fmt.Printf("iValue value is %v, type is %T\n", iValue, iValue)
	f, ok := iValue.(float64)
	if ok {
		fmt.Println("float64 f is ", f)
	}
}

func main() {
	var b float64 = 2.2
	testFloat(b)
}
