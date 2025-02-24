package leedcode



func singleNumber(nums []int) int {
	count := make(map[int]int)
	for _, value := range nums {
		count[value]++
	}
	for _, value := range nums {
		if(count[value] == 1){
			return value
		}
	}
	return 0
}
