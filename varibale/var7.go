package main
import "fmt"
func main() {
	var i = 2
	var j = 3
	n := i + j
	fmt.Println("n=", n)

	str := "luyuan"
	// num := 1
	str1 := "hello"
	// invalid operation: str + num (mismatched types string and int)
	// res := str + num
	res := str + str1
	fmt.Println("res=", res)
}
