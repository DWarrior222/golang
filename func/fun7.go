// 如果一个文件同时包含全局变量定义，init 函数和 main 函数，则执行的流程全局变量定义->init函数->main 函数
package main
import "fmt"

var age = test()

func test() int {
	fmt.Println("test函数")
	return 11
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main", age)
}
