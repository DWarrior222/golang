package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	_, err = conn.Do("Set", "name", "猫")
	if err != nil {
		fmt.Println("set err is ", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err is ", err)
		return
	}

	// nameString := r.(string)
	fmt.Println("操作ok", r)
}
