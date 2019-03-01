package main

import (
	"fmt"
	"os"
)

var userID int
var userPwd string

func main() {
	var key int
	var loop = true

	for loop {
		fmt.Println("---------------欢迎登录多人聊天系统---------------")
		fmt.Println("1 登录聊天系统\n2 注册用户\n3 退出系统")
		fmt.Println("-----------------------------------------------")
		fmt.Println("请选择(1-3):")

		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天系统")
			loop = false
		case 2:
			fmt.Println("注册")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}

		if key == 1 {
			fmt.Println("请输入用户id")
			fmt.Scanln(&userID)
			fmt.Println("请输入密码")
			fmt.Scanln(&userPwd)

			err := Login(userID, userPwd)
			if err != nil {
				fmt.Println("登录失败")
			} else {
				fmt.Println("登录成功")
			}
		} else {
			fmt.Println("用户注册的逻辑")
		}
	}
}
