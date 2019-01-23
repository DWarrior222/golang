package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度，按字节 len(str)
	str := "hello,北京"
	fmt.Println("str len=", len(str))

	// 字符串遍历，同时处理有中文的问题 r := []rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数
	// n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("结果是", n)
	}

	// 整数转字符串 str = strconv.Itoa(12345)
	str1 := strconv.Itoa(12345)
	fmt.Printf("str2=%v, str2=%T", str1, str1)
	// 10 进制转 2, 8, 16 进制:
	fmt.Println("123对应的二进制位", strconv.FormatInt(123, 2))
	// 查找子串是否在指定的字符串中
	fmt.Println("seafood里面有foo吗", strings.Contains("seafood", "foo"))

	// 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组: strings.Split("hello,wrold,ok", ",")
	fmt.Println("分隔字符串为数组hello,world,ok", strings.Split("hello,wrold,ok", ","))

	// 将字符串的字母进行大小写的转换
	fmt.Println("Go转化为大写", strings.ToLower("Go"))

	// 判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
	fmt.Println("判断字符串是否以指定的字符串开头", strings.HasPrefix("ftp://192.168.10.1", "ftp"))

	// 判断字符串是否以指定的字符串结束:
	fmt.Println("判断字符串是否以指定的字符串结束:", strings.HasSuffix("NLT_abc.jpg", "abc"))
}
