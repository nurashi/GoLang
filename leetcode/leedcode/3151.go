package leedcode

func isArraySpecial(nums []int) bool {

	for i := 0; i < len(nums) - 1; i++ {
		if (nums[i]%2 == 1 && nums[i+1]%2 == 1) || (nums[i]%2 == 0 && nums[i+1]%2 == 0) {
			return false
		}
	}
	return true
}
