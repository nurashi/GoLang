package leedcode

func IsPalindrome(word string) bool {

	for i := range len(word) / 2 {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func firstPalindrome(words []string) string {

	for i := 0; i < len(words); i++ {
		checker := IsPalindrome(words[i])
		if(checker) {
			return words[i] 
		}
	}
	return ""
}