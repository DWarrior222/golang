// 命令行参数
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	for k, v := range os.Args {
		fmt.Printf("args[%v]=%v \n", k, v)
	}

	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认3306")

	flag.Parse()

	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v \n", user, pwd, host, port)
}
