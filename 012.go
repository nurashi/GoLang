package main

import (
	"fmt"
)

func polindrom(a int) bool {
	if a == a[::-1] {
		return true
	}else {
		return false
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(polindrom(n))
}
