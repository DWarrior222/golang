package main
import "fmt"
func main() {
	var i int = 10
	fmt.Println("i的地址=", &i)

	// 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)

	var num int = 100
	var ptr1 *int = &num
	*ptr1 = 101
	fmt.Println("num=", num)
}
