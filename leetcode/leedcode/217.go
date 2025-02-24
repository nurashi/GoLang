package leedcode



func containsDuplicate(nums []int) bool {
	count := make(map[int]bool)
	for _, val := range nums {
		if count[val] {
			return true
		}
		count[val] = true
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	// sort.Ints(nums)


	for i := 0; i < len(nums) - 1; i++ {
		if(nums[i] == nums[i+1]) {
			return true
		}
	}
	return false
}