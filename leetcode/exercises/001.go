package main

import "fmt"

func Loop(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	Loop(n)
}
