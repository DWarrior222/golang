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
	var slice1 = arr[:len(arr)]
	var slice2 = arr[1:]
	var slice3 = arr[:]
	fmt.Println("arr=", arr)
	fmt.Printf("slice=%v, slice1=%v, slice2=%c, slice3=%v", slice, slice1, slice2, slice3)

	var slice4 []int = []int{11, 22, 33, 44, 55}
	slice4 = append(slice4, 66, 77)
	fmt.Println("slice4=", slice4)
	slice4 = append(slice4, slice4...)
	fmt.Println("slice4 double=", slice4)

	var slice5 = make([]int, 10)
	fmt.Println("slice5========", slice5)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)

	var slice6 []int

	var slice7 = []int{1, 2, 3, 4, 5}

	fmt.Printf("slice6=%v, slice7=%v\n", slice6, slice7)

	// 如果需要修改字符串，可以先将 string -> []byte / 或者 []rune -> 修改 -> 重写转成 string
	// 转成[]byte后，可以处理英文和数字，但不能处理中文，原因是[]byte按字节来处理，而一个汉字占用三个字节，所以会出现乱码
	// 解决办法是转化为[]rune即可，[]rune是按字符处理的，兼容汉字
	var str string = "hello，golang"
	var slice8 = str[:]
	var strbyte []byte = []byte(str)
	strbyte[0] = 's'
	str = string(strbyte)
	fmt.Printf("str地址是%p值是%v, strbyte是%v slice8是%v\n", &str, str, strbyte, slice8)


	str1 := "中国"
	var strrune []rune = []rune(str1)
	strrune[0] = '北'
	strrune[1] = '京'
	str1 = string(strrune)
	fmt.Printf("str1是%v", str1)
}
