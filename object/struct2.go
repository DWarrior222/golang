package main
import(
	"fmt"
	"encoding/json"	
)

type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main () {
	var cat1 Cat

	// var cat2 Cat = Cat{}
	cat2 := Cat{
		"cat2", 11, "white", "=.=",
	}
	// var cat3 *Cat = new(Cat)
	var cat3 *Cat = new(Cat)
	(*cat3).Name = "cat3"
	cat3.Name = "cat33"
	(*cat3).Age = 12

	var cat4 *Cat = &Cat{}

	fmt.Printf("cat1=%v, cat2=%v, cat3=%v, cat4=%v", cat1, cat2, cat3, cat4)

	// 第 3 种和第 4 种方式返回的是 结构体指针
	// 结构体指针访问字段的标准方式应该是:(*结构体指针).字段名 ，比如 (*person).Name = "tom"
	// 但go做了一个简化，也支持结构体指针.字段名, 比如person.Name="tom"。更加符合程序员 使用的习惯，go 编译器底层 对 person.Name 做了转化 (*person).Name。


	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	jsonStr, err := json.Marshal(monster)
	fmt.Println("jsonStr", jsonStr)

	if err != nil {
		fmt.Println("json 处理错误", err)
	}

	fmt.Println("jsonStr", string(jsonStr))
}
