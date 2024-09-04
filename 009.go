package main

import "fmt"

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(isEven(number))
}
