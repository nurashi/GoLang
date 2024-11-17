package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	var result[]int
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if(arr[i] > 0) {
			result = append(result, arr[i])
		}
	}
	for i := 0; i < len(result);i++ {
		fmt.Print(result[i], " ")
	}
}
	