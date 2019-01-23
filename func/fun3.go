package main
import "fmt"
func f(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2 * f(n - 1) + 1
	}
}
func main() {
	res := f(3)
	fmt.Println("res=", res)
	fmt.Println("res=", f(4))
	fmt.Println("res=", f(5))
	fmt.Println("res=", f(6))
}
