package leedcode

func differenceOfSums(n int, m int) int { 
	sum := 0
	sumM := 0
	for i := 0; i <= n; i++ {
		if(i % m == 0) {
			sumM += i
		} else {
			sum += i
		}
	}

	return sum - sumM
}
