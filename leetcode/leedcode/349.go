package leedcode

import "slices"


func intersection(nums1 []int, nums2 []int) []int {
	var arr []int

	for _, v := range nums1 {
		if slices.Contains(nums2, v) && !slices.Contains(arr, v) {
			arr = append(arr, v)

		}
	}

	return arr


}
