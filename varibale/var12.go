// 字符类型使用细节
package main
import "fmt"
func main() {
	// 在 Go 中，字符的本质是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值
	// 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的 unicode 字符
	var n = 22269
	// 字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码.
	var n1 = 10 + 'a'

	fmt.Printf("n=%c, n1=%d", n, n1)
}
