package main
import "fmt"
// import "unsafe"
func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}
func main() {
	sum, sub := getVal(30, 30)
	fmt.Println("sum=", sum, "subp=", sub)
	sum2, _ := getVal(10, 30)
	fmt.Println("sum2=", sum2)
}
