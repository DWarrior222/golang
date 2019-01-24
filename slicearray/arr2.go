package main
import(
	"fmt"
)

func main() {
	// 四种初始化数组的方式
	var intArr [3]int = [3]int{1, 2, 3}
	fmt.Println("intArr=", intArr)
	var intArr1 = [3]int{5, 6, 7}
	fmt.Println("intArr1=", intArr1)

	var intArr2 = [...]int{8, 9, 10}
	fmt.Println("intArr2=", intArr2)

	var intArr3 = [...]int{1: 800, 0: 900, 2:999}
	fmt.Println("intArr3=", intArr3)
}
