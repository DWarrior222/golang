package main
import "fmt"
func main() {
	var age int
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	// fmt.Scanf("%d", &age)
	if age > 18 {
		fmt.Println("你的年龄大于18，你要对自己负责")
	}

	if height := 181; height > 180 {
		fmt.Println("你高于180，你属于高个子")
	}
}
