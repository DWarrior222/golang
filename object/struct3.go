package main

import (
	"fmt"
)

type A struct {
	Name string
}

type Person struct {
	Name string
}

func (person Person) speak() {
	fmt.Printf("%v, 是个大好人", person.Name)
}

func (person Person) getSum(n1, n2 int) int {
	return n1 + n2
}
func (person Person) calcu() {
	res := 0
	for i := 0; i < 1000; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是", res)
}

func (a A) test() {
	fmt.Println(a.Name)
}
func main() {
	var a1 A
	a1.Name = "tom"
	a1.test()

	var person Person
	person.Name = "路远"
	person.speak()
	person.calcu()
	sum := person.getSum(1, 2)
	fmt.Println("sum", sum)
}
