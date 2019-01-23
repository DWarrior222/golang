package main
import "fmt"
func f(n int, args... int) int {
	sum := n
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
func main() {
	// 可变参数的应用
	sum := f(1, 2, 3, 4, 5, 6, 10)
	fmt.Println("sum=", sum)
}
