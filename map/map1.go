package main
import(
	"fmt"
)

func main() {
	var a map[string]string
	var b map[string]int
	var c map[int]string
	var d map[string]map[string]string
	fmt.Printf("a=%v,b=%v,c=%v,d=%v", a, b, c, d)

	// 方式一
	var e map[string]string
	e = make(map[string]string, 10)
	e["1"] = "宋江"
	fmt.Println("e=", e)

	// 方式二
	f := make(map[string]string)
	f["n1"] = "北京"
	fmt.Println("f=", f)

	// 方式三
	g := map[string]string{
		"1": "路远",
	}
	fmt.Println("g=", g)

	stu := make(map[string]map[string]string)

	stu["s1"] = make(map[string]string, 3)
	stu["s1"]["name"] = "luyuan"
	stu["s1"]["sex"] = "boy"

	stu["s2"] = make(map[string]string, 3)
	stu["s2"]["name"] = "hello"
	stu["s2"]["sex"] = "girl"

	fmt.Println(stu)

	stu["s2"]["name"] = "xiaoming"
	delete(stu, "s1")
	fmt.Println(stu)
	// 如果我们要删除 map 的所有 key ,没有一个专门的方法一次删除，可以遍历一下 key, 逐个删除 或者 map = make(...)，make 一个新的，让原来的成为垃圾，被 gc 回收

	val, ok := stu["s2"]
	if ok {
		fmt.Printf("ok = %v, val = %v", ok, val)
	} else {
		fmt.Println("没有s1")
	}
}
