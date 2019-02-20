package main

import (
	"encoding/json"
	"fmt"
)

type People1 struct {
	Name string
	Age  int
}

func main() {
	// struct反序列化
	str := "{\"Name\":\"关\",\"Age\":24}"

	people := People1{}
	err := json.Unmarshal([]byte(str), &people)
	if err != nil {
		fmt.Println("反序列化错误，err is ", err)
	}
	fmt.Printf("struct反序列化后的结果是%v，name=%v，age=%v\n", people, people.Name, people.Age)

	// map反序列化
	str1 := "{\"Name\":\"关\",\"Age\":24}"

	a := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(str1), &a)
	if err1 != nil {
		fmt.Println("反序列化错误，err is ", err1)
	}
	fmt.Println("map反序列化后的结果", a)

	// slice反序列化
	str2 := "[{\"Name\":\"关\",\"Age\":24}]"
	var b []map[string]interface{}

	err2 := json.Unmarshal([]byte(str2), &b)
	if err2 != nil {
		fmt.Println("反序列化错误，err is ", err2)
	}
	fmt.Println("slice反序列化后的结果", b)

}
