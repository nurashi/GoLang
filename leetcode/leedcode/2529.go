package leedcode

func maximumCount(nums []int) int {

	size := len(nums)
	counterMin, zero := 0, 0

	for i := range nums {
		if nums[i] < 0 {
			counterMin++
		} else if nums[i] == 0 {
			zero++
		}
	}

	if counterMin > size-zero-counterMin {
		return counterMin
	}
	return size - zero - counterMin
}
