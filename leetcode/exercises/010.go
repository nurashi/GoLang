package main

import "fmt"

func inputArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

func bubbleSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	inputArray(arr, n)
	bubbleSort(arr, n)
	if arr[n-1] == arr[n-2] {
		fmt.Print("Second is Not Found")
	} else {
		fmt.Print(arr[n-2])
	}
}
