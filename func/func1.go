package main
import "fmt"
func getSumAndSub(n1 int, n2 int) (int, int) {
	var sum int = n1 + n2
	var sub int = n1 - n2
	return sum, sub
}
func main() {
	res1, res2 := getSumAndSub(11, 1)
	fmt.Printf("res1 = %d, res2 = %d", res1, res2)
}
