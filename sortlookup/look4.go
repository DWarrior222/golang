package main

import (
	"fmt"
)

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, lookupNum int) {
	// fmt.Printf("%v, %v, %v, %v", arr, leftIndex, rightIndex, lookupNum)
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}

	findIndex := (leftIndex + rightIndex) / 2
	if lookupNum == (*arr)[findIndex] {
		fmt.Printf("找到了，下标为%v", findIndex)
		return
	} else if (*arr)[findIndex] < lookupNum {
		leftIndex = findIndex + 1
		BinaryFind(arr, leftIndex, rightIndex, lookupNum)
	} else if (*arr)[findIndex] > lookupNum {
		rightIndex = findIndex - 1
		BinaryFind(arr, leftIndex, rightIndex, lookupNum)
	}
}
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	var lookupNum int
	var leftIndex int
	rightIndex := len(arr) - 1
	fmt.Println("请输入要查找的数字", lookupNum)
	fmt.Scanln(&lookupNum)
	BinaryFind(&arr, leftIndex, rightIndex, lookupNum)
}
