package main
import "fmt"
func main() {
	// Go中，数据类型的转换可以是从表示范围小-->表示范围大，也可以范围大--->范围小
	// 被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化!
	// 在转换中，比如将 int64 转成 int8 【-128---127】 ，编译时不会报错，只是转换的结果是按 溢出处理，和我们希望的结果不一样。 因此在转换时，需要考虑范围.
	var a int64 = 1100
	var b float32 = float32(a)
	var c byte = byte(a)
	var e int8 = int8(a)
	// fmt.Printf("a的类型 %T a的占用字节数是%d", a, unsafe.Sizeof(a))
	fmt.Printf("a=%v,b=%v,c=%v,e=%v", a, b, c, e)
	fmt.Printf("a的数据类型%T\n", a)

	var n1 int32 = 12
	var n2 int8
	var n3 int8

	n2 = int8(n1) + 127
	// constant 128 overflows int8
	// n3 = int8(n1) + 128
	fmt.Println(n2, n3)
}
