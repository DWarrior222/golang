package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json: "people_name"`
	Age  int    `json: "people_age"`
}

func main() {
	// struct序列化
	people := People{
		Name: "关",
		Age:  24,
	}
	data, err := json.Marshal(&people)
	if err != nil {
		fmt.Println("序列化错误，err is ", err)
	}
	fmt.Println("struct序列化后的结果", data, "----", string(data))

	// map序列化
	// var a map[string]interface
	// a = make(map[string]interface{})
	a := make(map[string]interface{})
	a["name"] = "关"
	a["age"] = 24

	data1, err1 := json.Marshal(a)
	if err1 != nil {
		fmt.Println("序列化错误，err is ", err1)
	}
	fmt.Println("map序列化后的结果", data1, "----", string(data1))

	// slice序列化
	var slice []map[string]interface{}
	b := make(map[string]interface{})
	b["name"] = "关"
	b["age"] = 24
	slice = append(slice, b)
	data2, err2 := json.Marshal(slice)

	if err2 != nil {
		fmt.Println("序列化错误，err is ", err2)
	}
	fmt.Println("slice序列化后的结果", data2, "----", string(data2))

	// 基本数据类型序列化
	number := 234.13
	data3, err3 := json.Marshal(number)
	if err3 != nil {
		fmt.Println("序列化错误，err is ", err3)
	}
	fmt.Println("基本数据类型序列化后的结果", data3, "----", string(data3))
}
