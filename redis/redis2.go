package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("dial err is ", err)
		return
	}

	// conn.Do
	_, err = conn.Do("Hset", "gsj", "name", "guanshuaijie")
	if err != nil {
		fmt.Println("Hset err is ", err)
		return
	}

	_, err = conn.Do("Hset", "gsj", "age", "22")
	if err != nil {
		fmt.Println("Hset err is ", err)
		return
	}

	_, err = conn.Do("HMSet", "gsj", "sex", "boy", "job", "web coder", "address", "beijing")
	if err != nil {
		fmt.Println("HMSet err is ", err)
		return
	}

	r, err := redis.String(conn.Do("Hget", "gsj", "name"))
	if err != nil {
		fmt.Println("Hget err is ", err)
		return
	}

	fmt.Println("Hget ok, it's ", r)

	// redis.Strings
	r1, err := redis.Strings(conn.Do("HMGet", "gsj", "name", "job", "address"))
	if err != nil {
		fmt.Println("HMGet err is ", err)
		return
	}
	fmt.Printf("r1 type is %T \n", r1)
	for i, val := range r1 {
		fmt.Printf("r1[%v]=%v, type is %T \n", i, val, val)
	}
}
