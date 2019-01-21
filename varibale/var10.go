// 浮点型使用细节
package main
import "fmt"
import "unsafe"
func main() {
	// Golang浮点类型有固定的范围和字段长度，不收具体OS的影响
	// 浮点类型默认为float64类型

	var n1 = 5.12
	var n2 = .123
	fmt.Println("n1=", n1, "n2=", n2)

	var n3 float32 = 1.1
	var n4 float64 = 1.1

	fmt.Printf("n3的数据类型 %T, 占用的字节数 %d", n3, unsafe.Sizeof(n3)) 
	fmt.Printf("n4的数据类型 %T, 占用的字节数 %d", n4, unsafe.Sizeof(n4)) 

}
