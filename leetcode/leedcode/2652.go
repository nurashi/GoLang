package leedcode


func sumOfMultiples(n int) int {
	sum := 0
	i := 0
	for i <= n {
		if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
			sum += i
		}
		i++
	}
	return sum
}