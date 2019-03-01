package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.4.134:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("连接服务器成功...")
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}

		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}
}
