package main

import (
	"fmt"
)

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v \n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v \n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v \n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v \n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型，值是%v \n", index, x)
		case Person:
			fmt.Printf("第%v个参数是Person类型，值是%v \n", index, x)
		case *Person:
			fmt.Printf("第%v个参数是*Person类型，值是%v \n", index, x)
		default:
			fmt.Printf("第%v个参数是不确定类型，值是%v \n", index, x)
		}
	}
}

type Person struct {
	Name string
}

func main() {
	var a int = 1
	var b float32 = 1.1
	var c string = "a"
	var d bool = true
	person := Person{
		Name: "luyuan",
	}
	TypeJudge(a, b, c, d, person, &person)
}
