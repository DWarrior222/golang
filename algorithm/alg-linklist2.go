// 链表
package main

import "fmt"

type HeroList struct {
	no       int
	name     string
	nickname string
	pre      *HeroList
	next     *HeroList
}

func InsertHeroList(head *HeroList, newHeroNode *HeroList) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

func InsertHeroListByNo(head *HeroList, newHeroNode *HeroList) {
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
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

func DelHeroList(head *HeroList, no int) {
	temp := head
	flag := true
	if temp.next == nil {
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
		// 找到了
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
		fmt.Println("删除成功")
	} else {
		// 没找到
		fmt.Println("没有找到")
	}
}

func ListHeroList(head *HeroList) {
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
	return
}

func main() {
	head := &HeroList{}
	hero1 := &HeroList{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroList{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroList{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroList{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}

	InsertHeroListByNo(head, hero3)
	InsertHeroListByNo(head, hero2)
	InsertHeroListByNo(head, hero4)
	InsertHeroListByNo(head, hero1)
	InsertHeroListByNo(head, hero1)

	fmt.Println("删除前....")
	ListHeroList(head)

	DelHeroList(head, 2)

	fmt.Println("删除后....")
	ListHeroList(head)

	DelHeroList(head, 2)
}
