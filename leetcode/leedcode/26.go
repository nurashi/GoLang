package main

import "fmt"

func inputArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

func removeDuplicates(arr []int, n int) int {
	var idx int = 1
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			arr[idx] = arr[i]
			idx++
		}
	}
	return idx
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

func output(arr []int, n int, idx int) {
	for i := 0; i < idx; i++ {
		fmt.Print(arr[i], " ")
	}
	for j := idx; j < n; j++ {
		fmt.Print("_ ")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	inputArray(arr, n)
	bubbleSort(arr, n)
	output(arr, n, removeDuplicates(arr, n))

}
