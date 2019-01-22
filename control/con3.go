package main
import "fmt"
func main() {
	var second float64
	fmt.Println("请输入成绩")
	fmt.Scanln(&second)
	// fmt.Scanf("%d", &age)
	if second < 8 {
		var gender string
		fmt.Println("请输入性别")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入男子组决赛")
		} else {
			fmt.Println("进入女子组决赛")
		}
	} else {
		fmt.Println("out...")
	}
}
