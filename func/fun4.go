package main
import "fmt"
func f(n *int) {
	*n = *n + 1
	fmt.Println("f中n=", *n)
}
func main() {
	var n int = 3
	// f(n)
	f(&n) // 如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型)，可以传 入变量的地址&，函数内以指针的方式操作变量。从效果上看类似引用 。
	fmt.Println("n=", n)
}
