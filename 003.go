package main

import "fmt"

func main() {
	var num, num1 int
	fmt.Scan(&num, &num1)

	if num%num1 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
