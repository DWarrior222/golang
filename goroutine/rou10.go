package main

import(
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan<- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
			case v := <-intChan:
				fmt.Printf("在intChan读取的数据是%v \n", v)
				time.Sleep(time.Second)
			case v := <-stringChan:
				fmt.Printf("从stringChan读取的数据%s\n", v)
				time.Sleep(time.Second)
			default:
				fmt.Println("添加逻辑")
				time.Sleep(time.Second)
				return
		}
	}
}