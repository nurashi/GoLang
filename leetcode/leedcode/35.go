package leedcode


func searchInsert(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= target {
			return i
		}
	}
	return len(arr)
}


