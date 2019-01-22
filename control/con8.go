package main
import "fmt"
func main() {
	// 字符串遍历方式 1-传统方式
	var str string = "hello, 中国"
	for i := 0; i < len(str); i++ {
		fmt.Printf("打印了%c", str[i]) // 这样遍历的中文输出为乱码，原因是传统的
		//对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码是对应 3 个字节。
	}

	// 字符串遍历方式 2-for - range 
	// var str string = "hello, 中国"
	// 对应 for-range 遍历方式而言，是按照字符方式遍历。因此如果有字符串有中文，也是 ok
	for index, val := range str {
		fmt.Printf("index=%d, val=%c\n", index, val) 
	}

	// 如何解决 需要要将 str 转成 []rune 切片.=> 体验一把

	str2 := []rune(str)
	for j := 0; j < len(str2); j++ {
		fmt.Printf("%c\n", str2[j])
	}

}
