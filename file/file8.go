// (拷贝|copy|复制)文件
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(srcPath string, dstPath string) (written int64, err error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("open src is err ", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, errDst := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0666)
	if errDst != nil {
		fmt.Println("open dst is err ", errDst)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	srcPath := "./sky.jpeg"
	dstPath := "./sky1.jpeg"
	_, err := CopyFile(srcPath, dstPath)
	if err == nil {
		fmt.Println("copy success")
	} else {
		fmt.Println("copy faild")
	}
}
