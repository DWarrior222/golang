package main

import(
	"fmt"
	"time"
	"sync"
)

// 因为没有对全局变量 m 加锁，因此会出现资源争夺问题，代码会出现错误
// 解决方案：加入互斥锁
var(
	factorialMap = make(map[int]int, 10)
	lock sync.Mutex
)

func Fac(n int) {
	factorial := 1
	for i := n; i > 0; i-- {
		factorial = factorial * i
	}
	lock.Lock()
	factorialMap[n] = factorial
	lock.Unlock()
}

func FacMap(n int) {
	for i := n; i > 0; i-- {
		go Fac(i)
	}
}

func main() {
	// FacMap(10)
	for i := 200; i > 0; i-- {
		go Fac(i)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for i, v := range factorialMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()		
}