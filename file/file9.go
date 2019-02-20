// 统计英文、数字、空格和其他字符数量
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func main() {
	count := CharCount{}
	filePath := "./test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err is ", err)
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		strR := []byte(str)
		for _, v := range strR {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
