package main

import(
	"fmt"
)

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan<-i
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read data is ", v)
	}
	exitChan<-true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go WriteData(intChan)
	// go ReadData(intChan, exitChan)

	// 循环，不让主线程退出
	for {
		_, ok := <-exitChan 
		if !ok {
			break
		}
	}
}