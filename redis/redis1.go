package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Dial err is ", err)
		return
	}

	defer conn.Close()

	_, err = conn.Do("Set", "name", "路远")

	if err != nil {
		fmt.Println("set err is ", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("Get err is ", err)
		return
	}

	fmt.Println("Get value is ", r)
}
