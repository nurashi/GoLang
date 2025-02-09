package leedcode

func areOccurrencesEqual(s string) bool {
	count := make(map[string]int)

	for _, letter := range s {
		count[string(letter)]++
	}

	helper := count[string(s[0])]

	helper1 := false
	for _, letter := range s {
		if helper == count[string(letter)] {
			helper1 = true
		} else {
			helper1 = false
		}
		//fmt.Println(count[string(letter)])
	}

	return helper1
}
