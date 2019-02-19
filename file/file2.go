// 读取文件的内容并显示在终端(带缓冲区的方式)
package main
import(
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)

func main() {
	// 打开文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err is", err)
	}

	// 函数退出时及时关闭file
	defer file.Close()

	/*
	const(
	defaultBufSize = 4096
	)
	*/
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Println(str)
	}

	fmt.Println("over")


	// 读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中)，这种方式适用于文件 不大的情况。相关方法和函数(ioutil.ReadFile)
	file1 := "./test.txt"
	content, err1 := ioutil.ReadFile(file1)

	if err1 != nil {
		fmt.Println("read file1 err is", err1)
	}
	fmt.Printf("%v", string(content))
}
