package main

import (
	"fmt"
)

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNodeByNo(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.name = newCatNode.name
		head.no = newCatNode.no
		head.next = head
		fmt.Println("newCatNode 加入到环形链表")
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCatNode(head *CatNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("没有数据")
		return
	}

	for {
		fmt.Printf("[%v %s] \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode, no int) *CatNode {
	helper := head
	temp := head

	if temp.next == nil {
		fmt.Println("空链表不能删除")
		return head
	}

	if temp.next == head {
		if temp.no == no {
			temp.next = nil
		}
		return head
	}

	for {
		// helper是 head之前节点 的引用
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		// 当遍历到最后一个，此时最后一个没有比较，在下面
		if temp.next == head {
			break
		}
		// 遍历最后一个之前的
		// 如果当前的是目标节点
		if temp.no == no {
			// 如果当前是head节点
			if temp == head {
				// 把head的next赋给head，现在head就是 之前head的next，相当于删除了 之前的head
				head = head.next
			}
			// 然后把 之前的head节点的next 赋值给 之前的head节点 前面的节点
			helper.next = temp.next
			fmt.Printf("cat%d 删除成功\n", no)
			flag = false
			break
		}
		// helper 永远引用的是 temp前一个节点
		temp = temp.next
		helper = helper.next
	}

	if flag {
		// 处理刚才没有比较的 最后一个节点，如果是目标节点，则删除。
		if temp.no == no {
			helper.next = temp.next
			fmt.Println("删除成功")
		} else {
			fmt.Println("没找到..")
		}
	}
	return head
}

func main() {
	head := &CatNode{}
	// cat1 := &CatNode{}
	// for i := 0; i < 100; i++ {
	// 	cat1 = &CatNode{
	// 		no:   i,
	// 		name: "cat" + string(i),
	// 	}
	// 	InsertCatNodeByNo(head, cat1)
	// }
	cat1 := &CatNode{
		no:   1,
		name: "cat1111111",
	}
	cat2 := &CatNode{
		no:   2,
		name: "cat2222222",
	}
	cat3 := &CatNode{
		no:   3,
		name: "cat3333333",
	}
	InsertCatNodeByNo(head, cat1)
	InsertCatNodeByNo(head, cat2)
	InsertCatNodeByNo(head, cat3)
	ListCatNode(head)

	head = DelCatNode(head, 1)
	fmt.Println("删除后")
	ListCatNode(head)
}
