package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (m *Monkey) Climbing() {
	fmt.Println("生来会爬树")
}

type LittleMonkey struct {
	Monkey
}

func (m *LittleMonkey) Flying() {
	fmt.Println("学会了飞")
}

func (m *LittleMonkey) Swimming() {
	fmt.Println("学会了游泳")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.Climbing()
	monkey.Flying()
	monkey.Swimming()
}
