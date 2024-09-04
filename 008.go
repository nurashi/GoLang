package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}
func main() {
	var number1, number2 int
	fmt.Scan(&number1, &number2)
	fmt.Println(min(number1, number2))
}
