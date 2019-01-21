package main
import "fmt"
func main() {
	var n1 int8 = 127
	// constant 128 overflows int8
	// n1 = 128

	var n2 uint8 = 255
	// constant 256 overflows uint8
	// n2 =256

	fmt.Println("n1=", n1, "n2=", n2)

	var c byte = 255
	fmt.Println("c=", c)
}
