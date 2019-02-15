package utils

import (
	"fmt"
)

type FamilyAccount struct {
	command string
	balance float64
	choice  string
	loop    bool
	money   float64
	note    string
	details string
	flag    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		command: "",
		balance: 10000,
		choice:  "",
		loop:    true,
		money:   0.0,
		note:    "",
		details: "收支\t账户余额\t收支金额\t说明",
		flag:    false,
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("------------------------当前收支明细记录------------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) signOut() {
	for {
		fmt.Println("你确定要退出吗？y/n")
		fmt.Scanln(&this.choice)
		if this.choice == "y" || this.choice == "yes" {
			this.loop = false
			break
		} else if this.choice == "n" || this.choice == "no" {
			break
		}
		fmt.Println("输入有误")
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("------------------------家庭收支记录软件------------------------")
		fmt.Printf("1 收支明细\n2 登录收入\n3 登记支出\n4 退出\n \n请选择(1-4):")
		fmt.Scanln(&this.command)

		switch this.command {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.signOut()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !this.loop {
			fmt.Println("您退出家庭收支记录软件")
			break
		}
	}
}
