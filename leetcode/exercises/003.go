package main

import "fmt"

func main() { 
	var n int
	sum := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Print(sum)
}