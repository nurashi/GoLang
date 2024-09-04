package main

import (
	"fmt"
	"math"
)

func isPrime(a int) bool {
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var number int
	fmt.Scan(&number)
	if isPrime(number) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
