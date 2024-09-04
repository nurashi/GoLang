package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)
	if age < 18 {
		fmt.Println("small")
	} else if age >= 18 && age < 65 {
		fmt.Println("big, but not biggest")
	} else {
		fmt.Println("biggest")
	}
}
