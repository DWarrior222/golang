package model

import "fmt"

type account struct {
	accountNo string
	balance   float64
	passwd    string
}

func NewAccount(accountNo string, passwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不正确")
		return nil
	}

	if len(passwd) != 6 {
		fmt.Println("密码长度不正确")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额数目不正确")
		return nil
	}

	return &account{
		accountNo: accountNo,
		balance:   balance,
		passwd:    passwd,
	}
}

func (acc *account) SetBalance(passwd string, balance float64) {
	// 判断
	acc.balance = balance
}
