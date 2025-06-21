package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {

	arr := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && nums[i] <= len(nums) {
			arr = append(arr, nums[i])
		}
	}
	return arr
}

func main() {
	arr := []int{4,3,2,7,8,2,3,1}
	fmt.Print(findDisappearedNumbers(arr))
}