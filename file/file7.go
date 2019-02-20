// 将一个文件的内容，写入到另外一个文件
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	exist, errStatus := PathExists("./write.txt")
	fmt.Println(exist, errStatus, 1111)

	readPath := "./write.txt"
	content, err := ioutil.ReadFile(readPath)

	if err != nil {
		fmt.Println("open file err is ", err)
	}

	readPath1 := "./write1.txt"
	content1, err1 := ioutil.ReadFile(readPath1)

	if err1 != nil {
		fmt.Println("open file err is ", err1)
	}
	slice := append(content, content1...)

	writePath := "./write1.txt"
	err2 := ioutil.WriteFile(writePath, slice, 0666)
	if err2 != nil {
		fmt.Println("write file err is ", err2)
	}
}
