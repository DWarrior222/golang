//使用数组来模拟一个栈的实现
package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int
	Top    int
	arr    [5]int
}

func (stack *Stack) Push(val int) (err error) {
	if stack.Top == stack.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	stack.Top++
	stack.arr[stack.Top] = val
	return
}

func (stack *Stack) Pop() (val int, err error) {
	if stack.Top == -1 {
		fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}
	val = stack.arr[stack.Top]
	stack.Top--
	return val, nil
}

func (stack *Stack) List() {
	if stack.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%v \n", i, stack.arr[i])
	}
}
func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(111)
	stack.Push(222)
	stack.Push(333)
	stack.Push(444)
	stack.Push(555)
	stack.List()
	val, err := stack.Pop()
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	fmt.Println("val is ", val)
	stack.List()
}
