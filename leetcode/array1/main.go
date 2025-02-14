package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	sum := 0
	convertedNumber := strconv.Itoa(num)

	for i := range convertedNumber {
		sum += int(convertedNumber[i] - '0')
	}

	if sum > 9 {
		addDigits(sum)
		sum = 0
	}

	return sum
}

func main() {

	num := 38
	fmt.Println(addDigits(num))
}
