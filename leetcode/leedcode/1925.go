package leedcode

import "math"


func countTriples(n int) int {
	result := 0
	for n > 0 {
		for i := 1; i < n; i++ {
			tmp := (n*n) - (i*i)
			r := int(math.Sqrt(float64(tmp)))
			if r*r == tmp {
				result++
			}
		}
		n--
	}
	return result
}
