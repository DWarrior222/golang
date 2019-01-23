package main
import "fmt"
func f(n2 float64, n1 float64) float64 {
	fmt.Printf("n1 type 是 %T\n", n1)
	return n1 + n2
}
func main() {
	fmt.Println("sum=", f(1, 2))
}
