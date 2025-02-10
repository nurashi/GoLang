package leedcode

import "math"

func scoreOfString(s string) int {

	score := 0
	for i := 0; i < len(s)-1; i++ {
		difference := int(s[i]) - int(s[i+1])
		score += int(math.Abs(float64(difference)))
	}

	return score
}
