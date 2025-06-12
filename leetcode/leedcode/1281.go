package leedcode


import "strconv"

func subtractProductAndSum(n int) int {
	mul := 1
	sum := 0
	strN := strconv.Itoa(n)

	
	for i := 0; i < len(strN); i++ {
		digit := int(strN[i] - '0')
		mul *= digit
		sum += digit
	}

	return mul - sum
}