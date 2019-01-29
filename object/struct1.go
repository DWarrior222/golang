package main
import(
	"fmt"
)

type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}
func main () {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"

	fmt.Println("cat1=", cat1)

	fmt.Println("猫猫的信息如下:")
	fmt.Println("名称", cat1.Name)
	fmt.Println("年龄", cat1.Age)
	fmt.Println("颜色", cat1.Color)
	fmt.Println("hobby", cat1.Hobby)
}
