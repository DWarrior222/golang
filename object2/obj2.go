package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
}

type D struct {
	a A
}

func main() {
	var c C
	c.A.Name = "luyuan"
	fmt.Println(c)

	var d D
	d.a.Name = "luyuan"
}
