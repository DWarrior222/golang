package main
import "fmt"
func main() {
	var month int
	var age int
	var price float64
	price = 60.00
	fmt.Println("请输入月份和年龄，用空格隔开")
	fmt.Scanf("%d %d", &month, &age)
	if month > 4 && month < 10 {
		if age < 18 {
			fmt.Println("票价为", price / 1)
		} else if age > 60 {
			fmt.Println("票价为", price / 3)
		} else {
			fmt.Println("票价为", price)
		}
	} else {
		if age < 18 || age > 60 {
			fmt.Println("票价为", 20)
		} else {
			fmt.Println("票价为", 40)
		}
	}
}
