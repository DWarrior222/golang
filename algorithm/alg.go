package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Println("v2 is ", v2)
		}
	}

	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	var sparseArr []ValNode
	sparseArr = append(sparseArr, valNode)

	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println("细数数组是.....")

	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	var chessMapReceiver [11][11]int

	for i, valNode := range sparseArr {
		if i != 0 {
			chessMapReceiver[valNode.row][valNode.col] = valNode.val
		}
	}
	fmt.Println("恢复后的数据是....")

	for _, v := range chessMapReceiver {
		for _, v2 := range v {
			fmt.Println("v2 is ", v2)
		}
	}
}
