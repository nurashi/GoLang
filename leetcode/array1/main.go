package main

import (
	"fmt"
	"slices"
)


//[5 4 2 3] -> [3 2 5 4]
// first round alice removes 2, Bob removes 3, append -> 3 and then 2, alice, bob, bob, alice

// find min first time, and remove, but save in some varible ALICE 
// find min second time, and remove, but save in same varible BOB
// append BOB varible to result
// append ALICE variable to result
func numberGame(nums []int) []int {

	result := []int{}

	// 2 3 4 5 -> 3 2 5 4
	slices.Sort(nums)
	for i := 1; i < len(nums); i+=2 {
		result = append(result, nums[i])
		result = append(result, nums[i - 1])
	}

	
	
	return result
}

func main() {
	nums := []int{5,4,2,3}

	result := numberGame(nums)

	fmt.Print(result, " ")
}