package main

import "fmt"

func removeDuplicates(arr []int, n int, m int) []int {
	cheker := false
	result := []int{}
	for i := 0; i < n+m; i++ {
		cheker = false
		for j := 0; j < len(result); j++ {
			if arr[i] == arr[j] {
				cheker = true
				break
			}
		}
		if !cheker {
			result = append(result, arr[i])
		}
	}
	return result
}

func inputArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

func toOneArray(arr []int, arr1 []int, n int, m int, arr2 []int) []int{
	return append(arr1, arr2...)
}

func output(result []int){
	for i:=0; i < len(result); i++ {
		fmt.Print(result[i], " ")
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	arr1 := make([]int, m)
	arr2 := make([]int, n+m)
	inputArray(arr, n)
	inputArray(arr1, m)
	oneArray :=	toOneArray(arr, arr1, n, m, arr2)
	uniqArray := removeDuplicates(oneArray, n, m)
	output(uniqArray)
}
