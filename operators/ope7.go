// 键盘输入语句
package main
import "fmt"
func main() {
	var a int = 1 >> 2
	var b int = -1 >> 2
	var c int = 1 << 2
	var d int = -1 << 2
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println(2&3)
	fmt.Println(2|3)

	var e int = 2&3
	fmt.Println("e=", e)
}

