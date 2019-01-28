package main
import(
	"fmt"
)

func main() {
	students := make(map[string]Stu, 10)

	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 28, "上海"}

	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	for k, v := range students {
		fmt.Println("学生的编号是%v", k)
		fmt.Println("学生的名字是%v", v.Name)
	}
}
