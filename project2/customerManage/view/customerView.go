package main

import(
	"fmt"
	"../service"
	"../model"
)

type customerView struct {
	command string
	loop bool
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("------------客户列表-------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n------------客户列表完成------------\n\n")
}

func (this *customerView) add() {
	fmt.Println("-------------添加客户------------")
	fmt.Println("请输入姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("请输入电邮:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("------------添加完成------------")
	} else {
		fmt.Println("------------添加失败------------")
	}
}

func (this *customerView) del() {
	fmt.Println("-------------删除客户------------")
	fmt.Println("请输入待删除客户的编号（-1退出）：")
	delNum := -1
	fmt.Scanln(&delNum)
	if delNum == -1 {
		return
	}
	for {
		fmt.Println("确认是否删除（y/n）:")
		choice := ""
		fmt.Scanln(&choice)
		if choice == "n" || choice == "N" {
			// return
			break
		} else if choice == "y" || choice == "Y" {
			if this.customerService.Del(delNum) {
				fmt.Println("-------------删除成功------------")
			} else {
				fmt.Println("-------------删除失败------------")
			}
			break
		} else {
			fmt.Println("请正确输入命令")
		}
	}
}

func (this *customerView) mainMeau() {
	for {
		fmt.Println("--------------客户信息管理软件--------------")
		fmt.Println("1 添加客户\n2 修改客户\n3 删除客户\n4 客户列表\n5 退    出\n 请选择（1-5）:")
		this.command = ""
		fmt.Scanln(&this.command)
		switch this.command {
			case "1":
				this.add()
			case "2":
				// fmt.Println("修改客户")
				this.modify()
			case "3":
				this.del()
			case "4":
				this.list()
			case "5":
				this.exit()
			default:
				fmt.Println("请输入正确的命令")
		}

		if !this.loop {
			fmt.Println("您退出了客户信息管理系统")
			break
		}
	}
}

func (this *customerView) exit() {
	fmt.Println("确认退出？（y/n）:")
	isLoop := ""
	for {
		fmt.Scanln(&isLoop)
		if (isLoop == "Y" || isLoop == "y" || isLoop == "N" || isLoop == "n") {
			break
		}
		fmt.Println("输入有误，重新输入")
	}
	if isLoop == "Y" || isLoop == "y" {
		this.loop = false
	}
}

func (this *customerView) modify() {
	fmt.Println("-------------修改客户------------")
	fmt.Println("请输入待修改客户的编号（-1退出）：")
	delNum := -1
	fmt.Scanln(&delNum)
	if delNum == -1 {
		return
	}
	customers := this.customerService.List()
	// keys := [5]string{"Name", "Gender", "Age", "Phone", "Email"}
	// for _, v := range keys {
	// 	v := customers[v]
	// }
	name := customers[delNum-1].Name
	gender := customers[delNum-1].Gender
	age := customers[delNum-1].Age
	phone := customers[delNum-1].Phone
	email := customers[delNum-1].Email
	fmt.Printf("姓名(%v):", customers[delNum-1].Name)
	fmt.Scanln(&name)
	fmt.Printf("性别(%v):", customers[delNum-1].Name)
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v):", customers[delNum-1].Name)
	fmt.Scanln(&age)
	fmt.Printf("电话(%v):", customers[delNum-1].Name)
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v):", customers[delNum-1].Name)
	fmt.Scanln(&email)
	customers[delNum-1].Name = name
	customers[delNum-1].Gender = gender
	customers[delNum-1].Age = age
	customers[delNum-1].Phone = phone
	customers[delNum-1].Email = email
	fmt.Println(customers[delNum-1])
	// 获取该客户的信息
	// 用户输入
	// 修改该客户的信息
}

func main() {
	customerView := customerView{
		command: "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMeau()
}