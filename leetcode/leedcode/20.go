package leedcode

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i]+1 == s[j] || s[i]+2 == s[j] {
				return true
			}
		}
	}

	return false
}
