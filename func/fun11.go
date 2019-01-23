package main
import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("当前时间为%v, type为%T\n", now, now)
	// 方式 1: 就是使用 Printf 或者 SPrintf
	fmt.Printf("%v年%v月%v日 %v时%v分%v秒", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 方式 1: 就是使用 Printf 或者 SPrintf
	dateStr := fmt.Sprintf("%v年%v月%v日 %v时%v分%v秒", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)

	// 方式二: 使用 time.Format() 方法完成:
	fmt.Printf(now.Format("2006-01-02 15:00:00"))
}
