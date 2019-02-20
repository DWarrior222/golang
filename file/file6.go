// 打开一个存在的文件，将原来的内容读出显示在终端，并且追加 5 句"hello,北京!"
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "./write1.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
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
		fmt.Println(str)
	}

	str1 := "Hello, Beijing\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str1)
	}
	writer.Flush()
}
