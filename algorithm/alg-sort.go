// 查找排序

// 选择排序(select sorting)也是一种简单的排序方法。
// 它的基本思想是:第一次从 R[0]~R[n-1]中选 取最小值，与 R[0]交换，
// 第二次从 R[1]~R[n-1]中选取最小值，与 R[1]交换，第三次从 R[2]~R[n-1]中选 取最小值，与 R[2]交换，
// ...，第 i 次从 R[i-1]~R[n-1]中选取最小值，与 R[i-1]交换，..., 第 n-1 次从 R[n-2]~R[n-1]中选取最小值，
// 与 R[n-2]交换，总共通过 n-1 次，得到一个按排序码从小到大排列的有序 序列。
package main

import "fmt"

func LookupSort() {
	arr := []int{1, 536, 34, 65, 34, 1234, 67, 89, 23}
	length := len(arr)
	for j := 0; j < length-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < length; i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		fmt.Printf("第%d次 %v \n", j+1, arr)
	}
}

func main() {
	LookupSort()
}
