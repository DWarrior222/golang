// 打开一个存在的文件，在原来的内容追加内容 'ASDFGHJKL'
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./write1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err is ", err)
		return
	}
	defer file.Close()
	str := "ASDFGHJKL"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
