package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", myAdd(2, 3))
	fmt.Println("4 + 6 =", myAdd(4, 6))
	fmt.Println("3 + 9 =", myAdd(3, 9))
}

func myAdd(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

