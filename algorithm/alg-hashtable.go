// 哈希表	散列表
// 使用链表来实现哈希表, 该链表不带表头
package main

import (
	"errors"
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Find(no int) (cur *Emp, err error) {
	cur = this.Head
	flag := false
	for {
		if cur == nil {
			break
		}
		if cur.Id == no {
			flag = true
			break
		}
		cur = cur.Next
	}
	if !flag {
		return nil, errors.New("没有找到")
	}
	return cur, nil
}

func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil
	//如果当前的 EmpLink 就是一个空链表
	if cur == nil {
		this.Head = emp
		return
	}

	//如果不是一个空链表,给 emp 找到对应的位置并插入
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}

	pre.Next = emp
	emp.Next = cur
}

func (this *EmpLink) ShowList(no int) {
	if this.Head == nil {
		fmt.Println("链表为空")
		return
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表是%d，雇员id是%d，名字是%s \n", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}

func (this *HashTable) HashFun(id int) (val int) { // 如果返回值确定(形参)，则函数体内不需要重新命名
	val = id % 7
	return val
}

func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowList(i)
	}
}

func (this *HashTable) Find(no int) (emp *Emp) {
	linkNo := this.HashFun(no)
	emp, err := this.LinkArr[linkNo].Find(no)
	if err != nil {
		fmt.Println("err is ", err)
		return nil
	}
	return emp
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员姓名")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("请输入雇员id")
			fmt.Scanln(&id)
			emp := hashTable.Find(id)
			if emp != nil {
				fmt.Printf("雇员id是%v，雇员姓名%s \n", emp.Id, emp.Name)
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("命令错误")
		}
	}
}
