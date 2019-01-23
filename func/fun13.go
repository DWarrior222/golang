package main
import(
	"fmt"
	"time"
)

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err) // runtime error: integer divide by zero 
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 /num2
	fmt.Println("res=", res)
}

func main(){
	num := new(int)
	fmt.Printf("num的类型%T,num的值%v,num的地址%v,num指向的值%v\n", num, num, &num, *num)
	*num = 100
	fmt.Printf("num的类型%T,num的值%v,num的地址%v,num指向的值%v\n", num, num, &num, *num)

	test()
	fmt.Println("后面的代码")
	time.Sleep(time.Second)

}
