package main
import(
	"fmt"
	// "lib/factmode/factmode"
	factmode "./demo/factmode"
)

func main () {
	var stu = factmode.student{
		Name: "luyuan",
	}

	fmt.Println(stu)
}
