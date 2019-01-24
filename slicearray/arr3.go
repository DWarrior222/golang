package main
import(
	"fmt"
)

func main() {
	// 数组的遍历 
	xiyouji := [...]string{"唐僧", "孙悟空", "猪八戒", "沙僧"}
	for i, v := range xiyouji {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("xiyouji[%d]=%v", i, xiyouji[i])
	}

	for _, v := range xiyouji {
		fmt.Printf("元素的值为%v", v)
	}

}
