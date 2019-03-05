package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	// "js", "java", "golang", "css"
	// _, err = conn.Do("lpush", "books", "no1:宋江", 30, "no2:卢俊义", 28)

	_, err = conn.Do("lpush", "books", "js", "java", "golang", "css")

	if err != nil {
		fmt.Println("err is", err)
		return
	}
	for {
		r, err := redis.String(conn.Do("rpop", "books"))
		if err != nil {
			fmt.Println("err is", err)
			// return
			break
		}

		fmt.Println(r)
	}

}
