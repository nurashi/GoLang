package main

import "fmt"

   
func intersection(nums1 []int, nums2 []int) []int {
    var result []int
	for i := 0; i < len(nums1); i++ {
		for j := i; j < len(nums2); j++ {
			if(nums1[i] == nums2[j]){
				result = append(result, nums1[i])
			}
		}
	}
	return result;
}

func main() { 
	arr := []int{1,2,2,1}
	arr1 := []int{2,2}
	result := intersection(arr, arr1)

	fmt.Println(result)
}