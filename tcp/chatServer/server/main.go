package main

import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"../common/message"
	"io"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")

	// conn.Read 在conn没有被关闭的情况下，才会阻塞
	// 如果客户端关闭了conn，就不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	// fmt.Println("pkgLen is ", pkgLen, "buf is", buf)

	// 根据 pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	fmt.Println("测试333333333333")

	// 把pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err is ", err)
		return
	}
	return
}

// 根据客户端发送消息的种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			// 处理登录
			err = serverProcessLogin(conn, mes)
		case message.RegisterMesType:
			// 处理注册
		default:
			fmt.Println("消息类型不存在,无法处理")
	}
	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)

	if err != nil {
		fmt.Println("json.Unmarshal fail err is ", err)
		return
	}

	// 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	// 如果用户id为100，密码为123456，合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在"
	}

	// 将loginResMes序列化
	// data, err := json.Marshal(loginResMes)
	// if err != nil {
	// 	fmt.Println("json.Marshal fail", err)
	// 	return
	// }
	return
}

func process(conn net.Conn) {
	// 延时关闭conn
	defer conn.Close()

	// 循环客户端发送的消息
	for {
		// 这里读取数据包，直接封装成一个函数，readPkg，返回Message，error
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return
			} else {
				fmt.Println("readPkg err is ", err)
				return
			}
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
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
