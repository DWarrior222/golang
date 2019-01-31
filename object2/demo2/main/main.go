package main

import (
	"fmt"

	model "../model"
)

func main() {
	var person = model.NewAccount("luyuan", "123456", 22)
	fmt.Println(*person)
}
