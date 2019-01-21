// 字符类型
package main
import "fmt"
// import "unsafe"
func main() {
	//如果我们保存的字符在 ASCII 表的,比如[0-1, a-z,A-Z..]直接可以保存到 byte

	var c1 = 'a'
	var c2 = '0'
	fmt.Println("n1=", c1, "n2=", c2)
	fmt.Printf("c1=%c,c2=%c", c1, c2)

	var c3 = '北'
	var c4 int = '北'

	fmt.Printf("c3= %c, c3对应的码值 %d,数据类型为%T", c3, c3, c3) 
	fmt.Printf("c4= %c, c4对应的码值 %d,数据类型为%T", c4, c4, c4) 

}
