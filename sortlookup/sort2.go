// 冒泡排序
package main

import (
	"fmt"
)

const (
	LENGTH = 8
)

func BubbleSort(values []int) {
	flag := true
	var n int = 0
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				n++
				continue
			}
		}
		if flag {
			break
		}
	}
	fmt.Printf("方法二排序后%v，排序的次数%v\n", values, n)
}

func main() {
	var tmp int
	number := []int{95, 45, 15, 78, 84, 51, 24, 12}
	fmt.Printf("方法一排序前%v\n", number)
	var n int = 0
	for i := 0; i < LENGTH; i++ {
		for j := LENGTH - 1; j > i; j-- {
			if number[j] < number[j-1] {
				tmp = number[j-1]
				number[j-1] = number[j]
				number[j] = tmp
				n++
			}
		}
	}
	fmt.Printf("方法一排序后%v，排序的次数%v\n", number, n)

	number1 := []int{95, 45, 15, 78, 84, 51, 24, 12}
	fmt.Printf("方法二排序前%v\n", number1)
	BubbleSort(number1)
}
