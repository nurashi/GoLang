package main

import "fmt"


func mostFrequentEven(nums []int) []int { 


	evens := []int{}


	for i := 0; i < len(nums); i++ {
		if(nums[i] % 2 == 0) {
			evens = append(evens, nums[i])
		}
	}

	count := make(map[int]int)

	for _, value := range evens {
		count[value]++
		
	}
	max := -1
	repeatedEvens := []int{}
	for _, value := range evens {
		if(value > max){
			max = value
		}
	}

	for key, value := range evens {
		if(max == value) {
			repeatedEvens = append(repeatedEvens, evens[key])
		}
	}

	return repeatedEvens
}


func main() { 



	nums := []int{0,1,2,2,4,4,1}

	fmt.Println(mostFrequentEven(nums))
}