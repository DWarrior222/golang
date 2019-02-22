package main

import (
	"fmt"
	"strconv"
	"time"
)

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("11111111 hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("222222222 hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test1()
	go test2()
	for i := 1; i < 6; i++ {
		fmt.Println("main() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
