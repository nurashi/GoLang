package main

import "fmt"

func main() {
	var num, num1, result int
	fmt.Scan(&num, &num1)
	var operator string
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = num + num1
	case "-":
		result = num - num1
	case "*":
		result = num * num1
	case "/":
		if num1 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		result = num / num1
	default:
		fmt.Println("Invalid operator")
		return
	}
	fmt.Println(result)
}
