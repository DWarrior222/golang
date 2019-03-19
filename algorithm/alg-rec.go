package main

import (
	"fmt"
)

func SetWay(myMap *[8][7]int, i int, j int) bool {
	if myMap[4][4] == 2 { // 出口
		return true
	} else {
		if myMap[i][j] == 0 {
			myMap[i][j] = 2
			fmt.Printf("myMap[%v][%v]=%v \n", i, j, myMap[i][j])
			if SetWay(myMap, i+1, j) {
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else if SetWay(myMap, i-1, j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else {
				myMap[i][j] = 3
				fmt.Printf("myMap[%v][%v]=%v \n", i, j, myMap[i][j])
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	var myMap [8][7]int
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[1][2] = 1
	myMap[3][2] = 1
	myMap[4][2] = 1
	myMap[5][2] = 1
	myMap[6][2] = 1

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕")

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
