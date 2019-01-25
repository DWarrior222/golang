package main

import (
	"fmt"
)

func main() {
	var intArr [3]int
	intArr = [3]int{1, 2, 3}
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Printf("intArr地址为%p, intArr[0]地址为%p, intArr[1]地址为%p，intArr[2]地址为%p", &intArr, &intArr[0], &intArr[1], &intArr[2])
}
