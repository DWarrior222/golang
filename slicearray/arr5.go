package main
import(
	"fmt"
)


func main() {
	var arr [26]string
	for i := 0; i < 26; i++ {
		arr[i] = string('A' + byte(i))
		// fmt.Println("i=", string('A' + byte(i)))
		// fmt.Printf("A %T", string('A'))
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%v\n", arr[i])
	}
}
