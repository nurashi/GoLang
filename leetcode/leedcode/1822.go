package leedcode


func arraySign(nums []int) int {
	counter := 0

	for i := 0; i < len(nums); i++ {
		if(nums[i] == 0){
			return 0
		} else if (nums[i] < 0) {
			counter++
		}
	}

	if counter % 2 == 0 {
		return 1
	} else {
		return -1
	}
}