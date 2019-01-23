package main
import (
	"fmt"
	libs "lib"
)
func main() {
	fmt.Println("hello")
	var sum float64 = libs.Add(1.1, 2.2)
	fmt.Println(sum)
}
