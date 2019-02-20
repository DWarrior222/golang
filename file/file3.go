// 创建一个新文件，写入内容 5 句 "hello, Gardon"
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./write.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err is ", err)
	}
	defer file.Close()
	str := "hello, Gardon\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

	filePath1 := ""
	file1, err1 := os.OpenFile(filePath1, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println("open file err is ", err1)
	}
	str1 := "你好路远\n"

	writer1 := bufio.NewWriter(file1)
	for j := 0; j < 10; j++ {
		writer1.WriteString(str1)
	}
	writer1.Flush()
}
