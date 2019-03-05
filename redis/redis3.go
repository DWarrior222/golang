package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("dial err is ", err)
		return
	}

	// conn.Do
	_, err = conn.Do("MSet", "name", "guanshuaijie", "address", "北京朝阳区")
	if err != nil {
		fmt.Println("Hset err is ", err)
		return
	}

	r, err := redis.Strings(conn.Do("MGet", "name", "address"))
	for _, v := range r {
		fmt.Println(v)
	}

	_, err = conn.Do("Set", "name1", "路远")
	_, err = conn.Do("expire", "name1", 2)
	r1, err := redis.String(conn.Do("Get", "name1"))
	if err != nil {
		fmt.Println("Hset err is ", err)
		return
	}
	fmt.Println(r1)

	time.Sleep(time.Second * 3)

	r1, err = redis.String(conn.Do("Get", "name1"))
	fmt.Println(r1)
}
