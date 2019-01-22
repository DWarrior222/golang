package main
import "fmt"
func main() {
	var x interface{}
	var y = "10"
	x = y
	switch i := x.(type) {
		case nil:
			fmt.Printf("x的类型是%T", i)
		case int:
			fmt.Println("x的类型是int型")
		case float64:
			fmt.Println("x的类型是float64型")
		case func(int) float64:
			fmt.Println("x的类型是func型")
		case bool, string:
			fmt.Println("x的类型是bool或者string型")
		default:
			fmt.Println("未知类型")
	}
}
