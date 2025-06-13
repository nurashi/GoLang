package leedcode


func countSegments(s string) int {
	if len(s) == 0 { 
		return 0
	}

	counter := 0
	inWord := false

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if !inWord {
				counter++
				inWord = true
			}
		} else {
			inWord = false
		}
	}

	return counter
}