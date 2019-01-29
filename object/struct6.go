package main

import (
	"fmt"
)

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

// (p Person) , 则是值拷贝， 如果和指针类型，比如是 (p *Person) 则是地址拷贝。
// 如果不使用指针类型，相当于值复制，声明了一个新的stu变量
// 如果想引用stu，则使用指针类型
func (stu *Student) say() string {
	stu.name = "lu"
	fmt.Println("say", stu.name)
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v]id=[%v]score=[%v]", stu.name, stu.gender, stu.age, stu.id, stu.score)
	return infoStr
}

func (stu1 *Student) say1() {
	stu1.name = "xiaoming--name"
	fmt.Println("say1", stu1.name)
}

func main() {
	var stu = Student{
		name:   "luyuan",
		gender: "boy",
		age:    18,
		id:     1000,
		score:  100.00,
	}
	var stu1 = &Student{
		name:   "xiaoming",
		gender: "boy",
		age:    28,
		id:     100,
		score:  20.00,
	}
	fmt.Println(stu.say())
	fmt.Println("main", stu.name)
	stu1.say1()
	fmt.Println("main", stu1.name)
}
