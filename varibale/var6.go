// 声明变量的几种方式
package main
import "fmt"

func main() {
	// 第一种:指定变量类型，声明后若不赋值，使用默认值
	var i int
	i = 10
	
	// 第二种:根据值自行判定变量类型(类型推导)
	var num = 10.11

	// 第三种:省略 var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	name := "luyuan"

	// 多变量声明(在编程中，有时我们需要一次性声明多个变量，Golang 也提供这样的语法)
	n1, name1, n2 := 1, "luyuan", 3

	fmt.Println("i=", i, "num=", num, "name=", name, "n1=", n1, "n2=", n2, "name1=", name1)
	// 该区域的数据值可以在同一类型范围内不断变化(重点)
	var intnum int
	intnum = 10
	intnum = 20

	fmt.Println("intnum=", intnum)

	// intnum = 12.12

 	// 变量在同一个作用域(在一个函数或者在代码块)内不能重名
	var num2 int = 55
	// num2 := 1
	// no new variables on left side of :=
	fmt.Println("num2=", num2)

	// 变量=变量名+值+数据类型，这一点请大家注意，变量的三要素
	// Golang 的变量如果没有赋初值，编译器会使用默认值, 比如 int 默认值 0 小数默认为 0 string 默认值为空串
}
