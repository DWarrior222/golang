package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (queue *Queue) AddQueue(num int) (err error) {
	if queue.rear == queue.maxSize-1 {
		return errors.New("queue fill")
	}
	queue.rear++
	queue.array[queue.rear] = num
	return
}

func (queue *Queue) GetQueue() (val int, err error) {
	if queue.front == queue.rear {
		return -1, errors.New("queue empty")
	}
	queue.front++
	val = queue.array[queue.front]
	return val, err
}

func (queue *Queue) ShowQueue() {
	for i := queue.front + 1; i < queue.rear; i++ {
		fmt.Printf("array[%v] = %v \n", i, queue.array[i])
	}
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("输入 add 表示添加数据到队列")
		fmt.Println("输入 get 表示从队列获取数据")
		fmt.Println("输入 show 表示显示队列")
		fmt.Println("输入 exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("err is ", err)
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println("err is ", err)
			} else {
				fmt.Println("队列获取的数据是", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
