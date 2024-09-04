package main

import "fmt"

func factorial(a int) int {
	factior := 1
	for i := 1; i <= a; i++ {
		factior *= i
	}
	return factior
}

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(factorial(number))
}
