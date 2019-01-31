package main

import (
	"fmt"

	model "../person"
)

func main() {
	var person = model.NewPerson("路远")
	fmt.Println(*person)
	fmt.Println((*person).GetAge())
	(*person).SetAge(10)
	fmt.Println((*person).GetAge())
}
