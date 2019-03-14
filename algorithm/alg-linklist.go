package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = newHeroNode
}

func InsertHeroNodeByNo(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("添加重复")
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func DelHeroNode(head *HeroNode, no int) {
	temp := head
	flag := true
	if temp.next == nil {
		fmt.Println("没有数据")
		return
	}
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == no {
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("删除成功")
		temp.next = temp.next.next
	} else {
		fmt.Println("没有找到")
	}
}

func ListHreoNode(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("没有数据")
		return
	}

	for {
		fmt.Printf("[%v, %s, %s] \n", temp.no, temp.name, temp.nickname)
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}

	InsertHeroNodeByNo(head, hero3)
	InsertHeroNodeByNo(head, hero2)
	InsertHeroNodeByNo(head, hero4)
	InsertHeroNodeByNo(head, hero1)

	ListHreoNode(head)
	fmt.Println("删除前..........")
	DelHeroNode(head, 2)
	fmt.Println("删除后..........")
	ListHreoNode(head)
	DelHeroNode(head, 2)
}
