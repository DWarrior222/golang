// 冒泡排序
package main

import (
	"fmt"
)

func sort(arr *[8]int) {
	fmt.Println("排序前", *arr)
	temp := 0
	n := 0

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
				n++
			}
		}
	}

	fmt.Println("排序后", *arr)
	fmt.Println("次数", n)
}
func main() {
	var arr = [8]int{95, 45, 15, 78, 84, 51, 24, 12}

	sort(&arr)
}
