package main
import "fmt"
func main() {
	var classNum int = 3
	var stuNum int = 5
	var allScore float64
	for i := 1; i <= classNum; i++ {
		// var ave float64
		var classScore float64
		for j := 1; j <= stuNum; j++ {
			var score float64
			fmt.Printf("请输入%d班第%d个同学的成绩:", i, j)
			fmt.Scanln(&score)
			classScore += score
			if j == 5 {
				fmt.Printf("%d班的平均分数是%f\n", i, classScore / 5.0)
			}
		}
		allScore += classScore
		if i == 3 {
			fmt.Printf("所有班级的平均分是%f", allScore / float64(classNum * stuNum))
		}
	}
}
