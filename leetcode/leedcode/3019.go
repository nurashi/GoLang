package leedcode

func countKeyChanges(s string) int {
	counter := 0
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = runes[i] + 32 // to lower 
		}
	}

	for i := 0; i < len(runes) - 1; i++ {
		if runes[i] != runes[i+1] {
			counter++
		}
	}

	return counter
}