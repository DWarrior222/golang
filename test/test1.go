package main

import (
	"fmt"
)

func AddUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}
func main() {
	res := AddUpper(10)

	if res != 55 {
		fmt.Printf("addUpper错误，返回值=%v，期望值=%v", res, 55)
	} else {
		fmt.Printf("addUpper正确，返回值=%v，期望值=%v", res, 55)
	}
}
