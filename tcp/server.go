package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)

		//1. 等待客户端通过 conn 发送信息
		//2. 如果客户端没有 wrtie[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())

		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("客户端退出 err is ", err)
			return
		}
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err is ", err)
		return
	}
	defer listen.Close()

	// 等待客户端连接
	fmt.Println("等待客户端连接")
	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Accept() err is ", err)
		} else {
			fmt.Printf("Accept() suc con  is %v 客户端 ip is %v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
