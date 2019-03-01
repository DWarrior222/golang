package main

import (
	"fmt"
)

func Login(id int, pwd string) (err error) {
	fmt.Printf("userId is %v, userPwd is %v", id, pwd)
	return nil
}
