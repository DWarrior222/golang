// 算术运算符
package main
import "fmt"
func main() {
	// 运算的数都是整数，那么结果会去掉小数部分，保留整数
	fmt.Println(10/4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)
	//如果需要保留小数，则需要浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	fmt.Println("10%3=", 10 % 3)
	fmt.Println("10%-3=", 10 % -3)
}
