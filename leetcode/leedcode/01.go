package main

import "fmt"


func inputArray(arr []int, n int){
	for i:=0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

func twoSum(arr []int, target int) []int {
	result := make([]int, 2)
	for i:=0; i < len(arr); i++ {
		for j:=i + 1; j < len(arr); j++ {
			if(arr[i] + arr[j] == target){
				result[0] = i;
				result[1] = j; 
				return []int{i, j}
			}
		}
	}
	return nil
}



func main() {
	var n, target int
	fmt.Scan(&n, &target)
	arr := make([]int, n)
	inputArray(arr, n)
	result := twoSum(arr, target)

	fmt.Print(result[0], result[1]);
}
