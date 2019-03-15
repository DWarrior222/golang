package main

// 插入排序
import (
	"fmt"
)

func InsertSort() {
	arr := [9]int{1, 536, 34, 65, 34, 1234, 67, 89, 23}

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}

		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后 %v \n", i, arr)
	}
}

func main() {
	InsertSort()
}
