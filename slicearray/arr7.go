package main
import(
	"fmt"
//	"math/rand"
//	"time"
)
// 切片使用的方式
func main() {
	// 义一个切片，然后让切片去引用一个已经创建好的数组，比如前面的案例就是这样的。
	var arr = [5]int{1, 2, 3, 5, 6}
	var slice = arr[1:3]
	fmt.Println("arr=", arr)
	fmt.Println("slice=", slice)

	// 第二种方式:通过 make 来创建切片.
	// 直接定义切片时，[]int中[]里面不需要具体的值
	var slice1 []int = make([]int, 3, 10)
	slice1[1] = 10
	slice1[2] = 20
	fmt.Println(slice1)
	fmt.Println("slice1的size为", len(slice1))
	fmt.Println("slice1的cap为", cap(slice1))

	// 第 3 种方式:定义一个切片，直接就指定具体数组，使用原理类似 make 的方式
	// 类似var arr1 [3]int = [3]int{1, 1, 1}
	var slice2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)
	fmt.Println("slice2的size为", len(slice2))
	fmt.Println("slice2的cap为", cap(slice2))
}
