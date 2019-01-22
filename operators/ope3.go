// 逻辑运算符
package main
import "fmt"
func test () bool {
	fmt.Println("函数测试")
	return true
}
func main() {
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 || age < 10 {
		fmt.Println("ok2")
	}

	var i int = 10
	if i < 9 && test() {
		fmt.Println("ok3")
	}

	if i < 9 || test() {
		fmt.Println("ok4")
	}
}
