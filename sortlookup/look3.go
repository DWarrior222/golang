package main
import(
	"fmt"
)

func main() {
	names := [4]string{"七", "麦", "数", "据"}
	lookupWord := ""
	index := -1
	fmt.Println("请输入要查找的字", lookupWord)
	fmt.Scanln(&lookupWord)
	for i := 0; i < len(names); i++ {
		if lookupWord == names[i] {
			index = i
		}
	}

	if index != -1 {
		fmt.Printf("找到%v, 下标为%v \n", lookupWord, index)
	} else {
		fmt.Println("没有找到", lookupWord)
	}
}
