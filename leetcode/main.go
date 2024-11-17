package main

import "fmt"

func main() {
	var n int
	counter := 0
	fmt.Scan(&n)
	arr := make([]int, n)
	result := make([]int, counter)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if(arr[i] > 0) {
			counter++
			result[counter] = arr[i]
		}
	}
	for i := 0; i < counter;i++ {
		fmt.Print(result[i], " ")
	}
}
