package main
import(
	"fmt"
)

type Person struct {
	Name string
	Age int
	Address string
}

func main() {
	personChan := make(chan Person, 10)
	per1 := Person{"luyuan", 11, "北京"}
	per2 := Person{"guan", 12, "北京"}
	per3 := Person{"shuai", 13, "北京"}
	per4 := Person{"jie", 14, "北京"}

	personChan <- per1
	personChan <- per2
	personChan <- per3
	personChan <- per4

	fmt.Printf("personChan len is %v, personChan cap is %v \n", len(personChan), cap(personChan))

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(<-personChan)
	// }
	close(personChan)

	for v := range personChan {
		fmt.Println(v)
	}
}