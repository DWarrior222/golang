package main
import "fmt"
import "unsafe"
func main() {
	var i int
	fmt.Println(i)

	var j = 10.11
	fmt.Println(j)

	name := "tom"
	fmt.Println(name)

	var a int = 1
	fmt.Println(a)

	var b int8 = 127
	fmt.Println(b)

	// var c uint8 = 256
	// constant 256 overflows uint8
	var c uint16 = 256
	fmt.Println(c)

	var d int64 = 10
	fmt.Println("%T %d", d, unsafe.Sizeof(d))

	var e byte = 'a'
	fmt.Println(e)

	var f byte = 'A'
	fmt.Println(f)
}
