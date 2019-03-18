//使用数组来模拟一个栈的使用
package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack2 struct {
	MaxTop int
	Top    int
	arr    [5]int
}

func (stack *Stack2) Push(val int) (err error) {
	if stack.Top == stack.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	stack.Top++
	stack.arr[stack.Top] = val
	return
}

func (stack *Stack2) Pop() (val int, err error) {
	if stack.Top == -1 {
		fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}
	val = stack.arr[stack.Top]
	stack.Top--
	return val, nil
}

func (stack *Stack2) List() {
	if stack.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%v \n", i, stack.arr[i])
	}
}

func (stack *Stack2) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

func (stack *Stack2) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("不支持该运算符")
	}
	return res
}

func (stack *Stack2) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	numStack := &Stack2{
		MaxTop: 20,
		Top:    -1,
	}
	operStack := &Stack2{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "30+30*6-4-6"
	// 扫描exp
	index := 0
	// 定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) {
			//如果 operStack 是一个空栈， 直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				//如果发现 opertStack 栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级
				//，就从符号栈 pop 出，并从数栈也 pop 两个数，进行运算，运算后的结果再重新入栈
				//到数栈， 当前符号再入符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {
			keepNum += ch

			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		if index+1 == len(exp) {
			break
		}
		index++
	}

	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		fmt.Println(result, numStack, operStack)
		numStack.Push(result)
	}

	res, _ := numStack.Pop()
	fmt.Printf("%v=%v", exp, res)
}
