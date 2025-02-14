
package leedcode

func majorityElement(nums []int) int {

	count := make(map[int]int)
	max := -1
	most := 0
	for _, num := range nums {
		count[num]++
	}


	for key, value := range count {
		if(value > max){
			max = value
			most = key
		}
	}

	return most
}