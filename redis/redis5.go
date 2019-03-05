package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个 ip 的 redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "狗狗")
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	fmt.Println("r =", r)

	conn2 := pool.Get()
	defer conn2.Close()
	_, err = conn2.Do("Set", "name2", "猫猫")
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	r, err = redis.String(conn2.Do("Get", "name2"))
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	fmt.Println("r = ", r)
}
