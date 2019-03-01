package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据...")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read() err is ", err)
			return
		}
		fmt.Println("读到的buf是", buf[:4])
	}
}

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	if err != nil {
		fmt.Println("net.Listen err is ", err)
		return
	}

	// 监听成功，等待客户端连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err is ", err)
		}
		go process(conn)
	}
}
