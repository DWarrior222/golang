// 基本类型转 string 类型
package main
import "fmt"
import "strconv"
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string
	// fmt.Sprintf().. 会返回转换后的字符串
	str = fmt.Sprintf("%d", num1)
	fmt.Println("str=", str)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)
	
	// 方式 2:使用 strconv 包的函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T, str=%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T, str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T, str=%q\n", str, str)


	// string 类型转基本数据类型
	var str3 string = "true"
	var b3 bool
	b3 , _  = strconv.ParseBool(str3)
	fmt.Printf("b3 type %T b3= %v\n", b3, b3)

	var str4 string = "1234590"
	var n4 int64
	var n5 int
	n4, _ = strconv.ParseInt(str4, 10, 64)
	n5 = int(n4)

	fmt.Printf("n4 type %T n4=%v\n", n4, n4)
	fmt.Printf("n5 type %T n5=%v\n", n5, n5)

	var str5 string = "123.456"
	var f1 float64
	f1 , _  = strconv.ParseFloat(str5, 64)
	fmt.Printf("f1 type %T f1= %v\n", f1, f1)
	// fmt.Printf("_ type %T _= %v\n", _, _)

	var str6 string = "hello"
	var n6 int64
	var f2 float64
	var b6 bool
	b6, _ = strconv.ParseBool(str6)
	n6, _ = strconv.ParseInt(str6, 10, 64)
	f2, _ = strconv.ParseFloat(str6, 64)
	fmt.Printf("b6 type %T b6=%v\n", b6, b6)
	fmt.Printf("n6 type %T n6=%v\n", n6, n6)
	fmt.Printf("f2 type %T f2=%v\n", f2, f2)
 

}
