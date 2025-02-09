package leedcode

import "fmt"

func inputArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
}

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func output(nums []int, count int) {
	for i := 0; i < count; i++ {
		fmt.Print(nums[i], " ")
	}
	for i := count; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
}

func main() {
	val := 3
	arr := make([]int, 100)
	inputArray(arr)
	count := removeElement(arr, val)
	output(arr, count)
}
