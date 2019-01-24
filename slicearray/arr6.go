package main
import(
	"fmt"
	"math/rand"
	"time"
)


func main() {
	var arr [5]int
	lens := len(arr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lens; i++ {
		// fmt.Printf("type %T, value %v \n", rand.Intn(100), rand.Intn(100))
		arr[i] = rand.Intn(100)
	}
	fmt.Println("交换前打印", arr)

	var tem int = 0
	for i := 0; i < lens/2; i++ {
		tem = arr[lens-1-i]
		arr[lens-1-i] = arr[i]
		arr[i] = tem
	}

	fmt.Println("交换后打印", arr)
}
