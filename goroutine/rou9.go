package main

import(
	"fmt"
	"time"
)

func WriteNum(numChan chan int) {
	for i := 0; i < 8000; i++ {
		numChan<-i
	}
	close(numChan)
}

func ReadNum(numChan chan int, primeNumChan chan int, exitChan chan bool) {
	for {

		time.Sleep(time.Millisecond * 10)
		v, ok := <-numChan
		if !ok {
			break
		}
		// 定义是素数
		lock := true
		for i := 2; i < v; i++ {
			if v % i == 0 { // 不满足
				lock = false
				break
			}
		}

		if lock {
			primeNumChan<-v
		}
	}
	exitChan<-true
}

func main() {
	numChan := make(chan int, 1000)
	primeNumChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go WriteNum(numChan)
	for i := 0; i < 4; i++ {
		go ReadNum(numChan, primeNumChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeNumChan)
	}()

	for {
		v, ok := <-primeNumChan
		if !ok {
			break
		}
		fmt.Println("prime number is ", v)
	}

	fmt.Println("main over")
}