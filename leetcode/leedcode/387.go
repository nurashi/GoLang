package leedcode



func firstUniqChar(s string) int {
	hashTable := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		hashTable[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if hashTable[s[i]] == 1 {
			return i
		}
	}
	return -1
}