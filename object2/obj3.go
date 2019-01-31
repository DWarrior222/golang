package main

import (
	"fmt"
)

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var tv1 TV
	tv1 = TV{
		Goods{
			Name:  "电视机",
			Price: 23456.00,
		},
		Brand{
			Name:    "海尔",
			Address: "北京",
		},
	}
	fmt.Println(tv1)
	tv1.Brand.Name = "夏普"
	fmt.Println(tv1)

	var tv2 TV2
	tv2 = TV2{
		&Goods{
			Name:  "电视机",
			Price: 23456.00,
		},
		&Brand{
			Name:    "海尔",
			Address: "北京",
		},
	}

	fmt.Println(tv2)
	fmt.Println(*tv2.Goods)
	// (*tv2).Brand.Name = "夏普"
	// fmt.Println((*tv2))
}
