package main

import (
	"fmt"
	// "factmode"
	factmode "../fact1"
)

func main() {
	// var stu = factmode.Student{
	// 	Name: "luyuan",
	// }
	var stu = factmode.NewStudent("luyuan", 32.1)

	fmt.Println(*stu)
	fmt.Printf("name是%v,score是%v", stu.Name, stu.GetScore())
}
