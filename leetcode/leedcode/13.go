package leedcode

func romanInt(romanNumber string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	resultNum := 0
	n := len(romanNumber)

	for i := 0; i < n; i++ {
		current := romanValues[romanNumber[i]]

		if i < n-1 && current < romanValues[romanNumber[i+1]] {
			resultNum += romanValues[romanNumber[i+1]] - current
			i++
		} else {
			resultNum += current
		}
	}

	return resultNum
}
