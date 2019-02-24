package main

import(
	"fmt"
)

func main() {
	//演示一下管道的使用
	//1. 创建一个可以存放 3 个 int 类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan的值%v,intChan地址%p\n", intChan, &intChan)

	intChan<-10
	num := 81
	intChan<-num
	// intChan<-50
	intChan<-9

	fmt.Printf("channel len=%v, cap=%v \n", len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	fmt.Println("Num2=", num2)
	fmt.Printf("channel len=%v, cap=%v \n",len(intChan), cap(intChan))

	num3 := <-intChan
	num4 := <-intChan
	// num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4, len(intChan), cap(intChan))
}
