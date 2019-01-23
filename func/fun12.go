package main
import(
	"fmt"
	"time"
)
const(
	Nanosecond = 1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second = 1000 * Millisecond
	Minute = 60 * Second
	Hour = 60 * Minute
)
func main() {
	fmt.Println(100 * time.Millisecond)
	fmt.Println("Hour", Hour)
	start := time.Now().Unix()
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
		if i == 20 {
			break
		}
	}
	end := time.Now().Unix()
	fmt.Printf("执行循环耗费的时间为%v秒", end - start)
}
