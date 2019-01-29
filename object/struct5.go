package main

import (
	"fmt"
)

type Utils struct {
}

func (mu Utils) Trans(arr *[4][4]int) {
	fmt.Printf("%v, %v", arr, *arr)
	fmt.Println("转置前")
	for _, v1 := range *arr {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

	for k1, v1 := range *arr {
		for k2, _ := range v1 {
			if k2 >= k1 {
				(*arr)[k1][k2], (*arr)[k2][k1] = (*arr)[k2][k1], (*arr)[k1][k2]
			}
		}
	}

	fmt.Println("转置后")
	for _, v1 := range *arr {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
}

func (mu Utils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("是奇数")
	}
}
func main() {
	var mu Utils
	mu.JudgeNum(2)

	// arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	mu.Trans(&arr)
}
