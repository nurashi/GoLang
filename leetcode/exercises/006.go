package main

import "fmt"

func main() {
	var n, indexMin, indexMax int
	fmt.Scan(&n)
	min := 75153
	max := -65489
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] < min {
			min = arr[i]
			indexMin = i
		}
		if arr[i] > max {
			max = arr[i]
			indexMax = i
		}
	}
	temp := arr[indexMin]
	arr[indexMin] = arr[indexMax]
	arr[indexMax] = temp
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}
