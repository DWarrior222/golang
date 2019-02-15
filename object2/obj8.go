package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point

	var b Point
	// b = a.(Point)就是类型断言，表示判断a是否是指向Point类型的变量，如果是就转成Point类型并赋值给b，否则报错
	b = a.(Point)
	fmt.Println(b)

	var x interface{}
	var c float32 = 3124.11
	x = c
	d := x.(float32)

	fmt.Println(d)

	var x1 interface{}
	var c1 float32 = 3245.23
	x1 = c1

	// 在进行类型断言时，如果类型不匹配，就会报 panic, 因此进行类型断言时，要确保原来的空接口
	// 在进行断言时，带上检测机制，如果成功就 ok,否则也不要报 panic
	if d1, ok := x1.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("d1的类型时%T，值是%v", d1, d1)
	} else {
		fmt.Println("convert faild")
	}
	fmt.Println("继续执行")
}
