package main
import "fmt"
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("打印了", i)
	}

	var j int = 1 
	for j < 10 {
		fmt.Printf("这是第%d个j\n", j)
		j++
	}

	k := 1
	for {
		if k <= 10 {
			fmt.Printf("这是第%d个k\n", k)
		} else {
			break
		}
		k++
	}
}
