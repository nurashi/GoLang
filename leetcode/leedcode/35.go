package leedcode

import "fmt"

func inputArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

func searchInsert(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= target {
			return i
		}
	}
	return len(arr)
}

func main() {
	var n, target int
	fmt.Scan(&n, &target)
	arr := make([]int, n)
	inputArray(arr, n)
	fmt.Print(searchInsert(arr, target))
}
