package leedcode


func twoSum(arr []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result[0] = i
				result[1] = j
				return []int{i, j}
			}
		}
	}
	return nil
}

