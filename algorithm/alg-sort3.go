// 快速排序
package main

import (
	"fmt"
)

//1. left 表示 数组左边的下标
//2. right 表示数组右边的下标
//3 array 表示要排序的数组
func QuickSort(left int, right int, array *[9]int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	fmt.Printf("left是 %v 中间值pivot是 %v right是 %v \n", l, pivot, r)
	temp := 0

	for l < r {
		for array[l] < pivot {
			l++
		}

		for array[r] > pivot {
			r--
		}

		if l >= r {
			break
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp

		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 下一次避开上一次中间的值，因为只需要把两边排序
	if l == r {
		l++
		r--
	}
	fmt.Printf("array = %v, left = %v, right = %v, l = %v, r = %v \n", array, left, right, l, r)
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	arr := [9]int{-9, 78, 1, 23, -567, 70, 123, 90, -23}
	fmt.Println("init array is", arr)
	QuickSort(0, len(arr)-1, &arr)
	// -9, 78, 1, 23, -567, 70, 123, 90, -23
	// left是 0 中间值pivot是 -567 right是 8
	// array = &[-567 78 1 23 -9 70 123 90 -23], left = 0, right = 8, l = 1, r = -1
	// left是 1 中间值pivot是 -9 right是 8
	// array = &[-567 -23 -9 23 1 70 123 90 78], left = 1, right = 8, l = 3, r = 1
	// left是 3 中间值pivot是 70 right是 8
	// array = &[-567 -23 -9 23 1 70 123 90 78], left = 3, right = 8, l = 6, r = 4
	// left是 3 中间值pivot是 23 right是 4
	// array = &[-567 -23 -9 1 23 70 123 90 78], left = 3, right = 4, l = 5, r = 3
	// left是 6 中间值pivot是 90 right是 8
	// array = &[-567 -23 -9 1 23 70 78 90 123], left = 6, right = 8, l = 8, r = 6
	fmt.Println("after sort is", arr)
}
