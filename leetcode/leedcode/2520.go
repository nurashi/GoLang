package leedcode


import "strconv"


func countDigits(num int) int {
	counter := 0
	strNum := strconv.Itoa(num)

	for i := 0; i < len(strNum); i++ {
		digit := int(strNum[i] - '0')
		if(num % digit == 0){ 
			counter++
		}
	}
	return counter
}