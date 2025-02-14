package leedcode

import (
	"math"
	"strconv"
)

func differenceOfSum(nums []int) int {

	sumElements := 0
	sumDigits := 0
	for _, number := range nums {
		sumElements += number

		if number > 9 {
			convertNumber := strconv.Itoa(number)
			for i := 0; i < len(convertNumber); i++ {
				sumDigits += int(convertNumber[i] - '0')
			}
		} else {
			sumDigits += number
		}
	}

	return int(math.Abs(float64(sumElements - sumDigits)))
}
