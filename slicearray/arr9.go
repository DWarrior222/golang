package main
import(
	"fmt"
)

func fbn(n int) int {
	if n < 2 {
		return 1
	}
	return fbn(n - 1) + fbn(n - 2)
}

func main() {
	var n int
	fmt.Scanln(&n)
	var slice []int = make([]int, n)
	// var slice []int
	for i := 0; i < n; i++ {
		slice[i] = fbn(i)
	}
	fmt.Println(slice)
}
