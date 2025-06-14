package leedcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashTable := make(map[byte]int)

	for i := 0; i < len(s); i++ { 
		hashTable[s[i]]++
		hashTable[t[i]]--
	}

	for _, value := range hashTable {
		if value != 0 {
			return false
		}
	}
	return true
}