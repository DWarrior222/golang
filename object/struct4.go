package main
import("fmt")

type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {
	stu.Name = "luyuan"
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main () {
	var stu Student
	stu.Name = "tom"
	stu.Age = 20
	fmt.Println(&stu)
}
