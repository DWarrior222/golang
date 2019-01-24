package main
import(
	"fmt"
)

func test(arr *[3]int) {
	(*arr)[0] = 88
	fmt.Printf("原数组地址%p, 修改后的原数组%v", arr, *arr)
}

func test1(arr []int) {
	fmt.Println("test1", arr)
}

func main() {
	arr := [3]int{1, 2, 3}
	// 如想在其它函数中，去修改原来的数组，可以使用引用传递(指针方式)
	test(&arr)
	// 长度是数组类型的一部分，在传递函数参数时 需要考虑数组的长度
	// test1(arr)
	test1([]int{})
	// invalid array index 3 (out of bounds for 3-element array)
	// fmt.Println(arr[3])
}
