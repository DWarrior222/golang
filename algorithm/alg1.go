// 稀疏数组
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ValNode struct {
	Row int `json:"row"`
	Col int `json:"col"`
	Val int `json:"val"`
}

func save() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	var chessSlice []ValNode

	valNode := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}

	chessSlice = append(chessSlice, valNode)

	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode := ValNode{
					Row: i,
					Col: j,
					Val: v2,
				}
				chessSlice = append(chessSlice, valNode)
			}
		}
	}

	fmt.Println(chessSlice)
	data, err := json.Marshal(chessSlice)

	if err != nil {
		fmt.Println("序列化错误，err is ", err)
	}
	fmt.Println("slice序列化后的结果", data, "----", string(data))

	filePath := "./chessmap.data"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// for _, val := range chessSlice {
	// 	fmt.Println("val is", val)
	// 	data1, err := json.Marshal(val)
	// 	if err != nil {
	// 		fmt.Println("err is ", err)
	// 		break
	// 	}
	// 	fmt.Println("struct序列化后的结果", data1, "----", string(data1))
	// 	writer.WriteString(string(data1))
	// }
	writer.WriteString(string(data))

	writer.Flush()
}

func get() {
	filePath := "./chessmap.data"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err == io.EOF {
		fmt.Printf("str is %v err is %v", str, err)
	}

	var chessSlice []ValNode
	err = json.Unmarshal([]byte(string(str)), &chessSlice)

	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	fmt.Println("slice反序列化结果", chessSlice)
	// 怎么设置一个不确定长度的数组
	var chessMap [11][11]int
	// chessMap := make(map[int]map[int]int)
	for i, v := range chessSlice {
		if i != 0 {
			chessMap[v.Row][v.Col] = v.Val
		}
	}
	fmt.Println(chessMap)
}

func main() {
	save()
	get()
}
