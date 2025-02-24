package main

import "fmt"
func maxProductDifference(nums []int) int {
//	sort.Ints(nums)
//	return nums[len(nums)-1]*nums[len(nums)-2] - (nums[0] * nums[1])
	var tempMax int
	var tempMin int
	max := -1
	min := 50000
	for _, i := range nums {
		if i > max {
			tempMax = max
			max = i
		}
		if i < min {
			tempMin = min
			min = i
		}
	}
	return tempMax*max - tempMin*min
}

func main() {
	nums := []int{5, 6, 2, 7, 4}
	fmt.Println(maxProductDifference(nums))
}
