package tasks



func factorial(n int) int {

	if(n > 1) {
		return n * factorial(n - 1) 
	} else {
		return 1
	}
}