// 整数的使用细节
package main
import "fmt"
import "unsafe"
func main() {
	//  Golang 各整数类型分:有符号和无符号，int uint 的大小和系统有关。
	// Golang 的整型默认声明为 int 型
	var n1 = 100

	fmt.Printf("n1的数据类型 %T \n", n1)

	// 如何在程序查看某个变量的字节大小和数据类型 (使用较多)
	var n2 = 1100
	var n3 int64 = 10
	fmt.Printf("n2的数据类型 %T, 占用的字节数 %d\n", n2, unsafe.Sizeof(n2)) 
	fmt.Printf("n3的数据类型 %T, 占用的字节数 %d", n3, unsafe.Sizeof(n3)) 

	// Golang程序中整型变量在使用时，遵守保小不保大的原则，即:在保证程序正确运行下，尽量 使用占用空间小的数据类型
	var age byte = 90

	// bit: 计算机中的最小存储单位。byte:计算机中基本存储单元。[二进制再详细说] 1byte = 8 bit

}
