// 覆盖 打开一个存在的文件中，将原来的内容覆盖成新的内容 10 句 "路远同学!"
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./write1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err is ", err)
		return
	}
	defer file.Close()
	str := "路远同学\r\n"
	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
