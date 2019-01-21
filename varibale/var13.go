package main
import "fmt"
// import "unsafe"
func main() {
	var a int
	var b float32
	var c byte
	var d bool
	var e string
	// fmt.Printf("a的类型 %T a的占用字节数是%d", a, unsafe.Sizeof(a))
	fmt.Printf("a=%d,b=%v,c=%v,d=%v,e=%v", a, b, c, d, e)
}
