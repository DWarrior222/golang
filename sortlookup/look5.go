package main
import(
	"fmt"
)

func main() {
	var arr [4][6]int

	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println(arr)
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	for m, v1 := range arr {
		for n, v2 := range v1 {
			fmt.Printf("arr[%v][%v]=%v\n", m, n, v2)
		}
	}
}
