// 队列
package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (queue *CircleQueue) Pop() (val int, err error) {
	if queue.IsEmpty() {
		return 0, errors.New("队列为空")
	}
	val = queue.array[queue.head]
	queue.head = (queue.head + 1) % queue.maxSize
	fmt.Println("queue.head =", queue.head)
	return
}

func (queue *CircleQueue) Push(val int) (err error) {
	if queue.IsFull() {
		return errors.New("队列满了")
	}
	queue.array[queue.tail] = val
	queue.tail = (queue.tail + 1) % queue.maxSize
	fmt.Println("queue.tail =", queue.tail)
	return
}
func (queue *CircleQueue) IsFull() bool {
	return (queue.tail+1)%queue.maxSize == queue.head
}

func (queue *CircleQueue) IsEmpty() bool {
	return queue.tail == queue.head
}

func (queue *CircleQueue) ShowQueue() {
	size := queue.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	head := queue.head
	for i := 0; i < size; i++ {
		fmt.Printf("%d. array[%d]=%d \n", i, head, queue.array[head])
		head = (head + 1) % queue.maxSize
	}
}

func (queue *CircleQueue) Size() (val int) {
	return (queue.tail + queue.maxSize - queue.head) % queue.maxSize
}
func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
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
		case "a":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("err is ", err)
			} else {
				fmt.Println("加入队列成功")
			}
		case "g":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println("err is ", err)
			} else {
				fmt.Println("队列获取的数据是", val)
			}
		case "s":
			queue.ShowQueue()
		case "e":
			os.Exit(0)
		}
	}
}
