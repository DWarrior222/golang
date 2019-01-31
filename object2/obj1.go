package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money < 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")

}

func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money < 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为%v，你的余额是%v", account.AccountNo, account.Balance)
}

func main() {
	var account = Account{
		AccountNo: "shuai0810",
		Pwd:       "shuai2121",
		Balance:   100.0,
	}
	var num int
	var userPwd string
	var userMoney float64
	fmt.Println("请输入以下选项的序号：\n1.查询余额\n2.存款\n3.取款")
	fmt.Scanln(&num)
	if num == 1 {
		fmt.Println("请输入你的密码")
		fmt.Scanln(&userPwd)
		account.Query(userPwd)
	} else if num == 2 {
		fmt.Println("请输入你的密码和存入的钱,用空格隔开")
		fmt.Scanf("%v %v", &userPwd, &userMoney)
		account.Deposite(userMoney, userPwd)
	} else {
		fmt.Println("请输入你的密码和取出的钱,用空格隔开")
		fmt.Scanf("%v %v", &userPwd, &userMoney)
		account.WithDraw(userMoney, userPwd)
	}
}
