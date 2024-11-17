package main

import "fmt"

func input(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

func output(result[] int, n int) {
	for i := n - 1; i >= 0; i-- {
		fmt.Print(result[i], " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	input(arr, n)
	output(arr, n)
}
