package main

import (
	"fmt"
	"encoding/json"
	"encoding/binary"
	"net"
	"../common/message"
)

func login(userId int,userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err is", err)
		return
	}

	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)

	if err != nil {
		fmt.Println("json.Marshal err is", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err is ", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}

	mes, err = readPkg(conn)

	if err != nil {
		fmt.Println("readPkg err is ", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
