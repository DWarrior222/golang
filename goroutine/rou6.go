package main

import(
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan<-10
	intChan<-20
	intChan<-30

	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan

	fmt.Printf("num1=%v,num2=%v,num3=%v \n", num1, num2, num3)

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["c1"] = "北京"
	m1["c2"] = "上海"
	m1["c3"] = "深圳"
	m2 := make(map[string]string, 20)
	m2["p1"] = "关"
	m2["p2"] = "帅"
	m2["p3"] = "杰"
	mapChan<-m1
	mapChan<-m2

	m4 := <-mapChan

	fmt.Printf("m4=%v,mapChan len = %v \n", m4, len(mapChan))

	var catChan chan Cat
	catChan = make(chan Cat, 10)
	cat1 := Cat{"kim", 1}
	cat2 := Cat{"kit", 2}
	catChan<-cat1
	catChan<-cat2

	cat3 := <-catChan
	fmt.Printf("cat3=%v,catChan len = %v \n", cat3, len(catChan))

	var catChan2 chan *Cat
	catChan2 = make(chan *Cat, 10)
	cat21 := &Cat{"kim", 1}
	cat22 := &Cat{"kit", 2}
	catChan2<-cat21
	catChan2<-cat22

	cat23 := <-catChan2
	fmt.Printf("cat3=%v,catChan len = %v \n", cat23, len(catChan2))

	// var allChan chan interface{}
	allChan := make(chan interface{}, 10)
	allChan <- cat1
	allChan <- 1
	allChan <- cat2
	allChan <- "cat1"

	all1 := <- allChan
	all2 := <- allChan
	all3 := <- allChan
	fmt.Printf("all1=%v,all2=%v,all3=%v,allChan len = %v,type = %T \n", all1, all2, all3, len(catChan2), all1)

	fmt.Printf("all1 type is %T，all1 value is %v \n", all1, all1)
	a := all1.(Cat)
	fmt.Println(a.Name)
}