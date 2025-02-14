package leedcode


func majorityElementMedium(nums []int) []int {
	n := len(nums)
	if(n < 3){
		return nums
	}
	

	count := make(map[int]int)
	max := -1
	for _, num := range nums {
		count[num]++
	}

	for _, value := range count {
		if(value > max){
			max = value
		}
	}
	arr := []int{}


	for key, value := range count {
		if(value > (n/3)) {
			arr = append(arr, key)
		}
	}

	return arr
}
