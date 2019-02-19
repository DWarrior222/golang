package main
import(
	"fmt"
	"os"
)

func main() {
	// file 叫 file对象
	// file 叫 file指针
	// file 叫 file文件句柄
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err is", err)
	}
	fmt.Printf("file=%v", file)

	err = file.Close()

	if err != nil {
		fmt.Println("close file err is", err)
	}
}
