// 如果一个文件同时包含全局变量定义，init 函数和 main 函数，则执行的流程全局变量定义->init函数->main 函数
package main
import "fmt"
// 定义全局变量
var (
	Fun1 = func (n1, n2 int) int {
		return n1 + n2
	}
)
func main() {

	res1 := func (n1, n2 int) int {
		return n1 + n2
	}(1, 2)

	fmt.Println("res1=", res1)

	a := func (n1, n2 int) int {
		return n1 + n2
	}

	res2 := a(1, 3)
	fmt.Println("res2=", res2)

	res3 := Fun1(1, 4)
	fmt.Println("res3=", res3)
}
