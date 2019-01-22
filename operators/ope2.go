// 算术运算符
package main
import "fmt"
func main() {
	var days int = 97
	var weeks int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期，零%d天\n", weeks, day)

	var huashi float32 = 123.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("%v 对应的摄氏温度=%v \n", huashi, sheshi)
}
