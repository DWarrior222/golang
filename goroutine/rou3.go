package main

import(
	"fmt"
)

func Factorial(n int) int {
	factorial := 1
	for i := n; i > 0; i-- {
		factorial = factorial * i
	}
	return factorial
}

func FactorialSum(n int) map[int]int {
	// factorialSum := 0
	factorialMap := make(map[int]int)
	for i := n; i > 0; i-- {
		// factorialSum += Factorial(i)
		factorialMap[i] = Factorial(i)
	}
	// return factorialSum
	return factorialMap
}

func main() {
	fmt.Println(FactorialSum(10))
}