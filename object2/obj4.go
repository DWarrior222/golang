package main

import (
	"fmt"
)

// Usb接口定义了两个方法
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

//让 Phone 实现 Usb 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机结束工作")
}

type Camera struct {
}

//让 Camera 实现 Usb 接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机结束工作")
}

type Computer struct {
}

// Computer结构体的Working方法时，接收一个 Usb 接口类型变量
// 只要是实现了 Usb 接口声明的所有方法，就是一个Usb接口类型变量
// usb根据传入的实参，来确定是哪个结构体
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	camera := Camera{}
	phone := Phone{}

	computer.Working(phone)
	computer.Working(camera)
}
