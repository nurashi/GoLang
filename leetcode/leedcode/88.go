package leedcode

import "fmt"

func inputNum1(nums1 []int, m int) {
	for i := 0; i < m; i++ {
		fmt.Scan(&nums1[i])
	}
}
func inputNum2(nums2 []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&nums2[i])
	}
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < n+m; i++ {
		for j := 0; j < n; j++ {
			nums1[i] = nums2[j]
		}
	}
	for i := 0; i < n+m-1; i++ {
		for j := 0; j < n+m-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				temp := nums1[j]
				nums1[j] = nums1[j+1]
				nums1[j+1] = temp
			}
		}
	}
}

func output(nums1 []int, n int, m int) {
	for i := 0; i < n+m; i++ {
		fmt.Print(nums1[i], " ")
	}
}

func main() {
	var n, m int
	fmt.Scan(&m, &n)

	nums1 := make([]int, m)
	nums2 := make([]int, n)
	inputNum1(nums1, m)
	inputNum2(nums2, n)
	merge(nums1, m, nums2, n)
	output(nums1, n, m)
}
