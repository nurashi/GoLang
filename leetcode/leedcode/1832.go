package leedcode


func checkIfPangram(sentence string) bool {
	
	var cheker [26]bool
	
	for i := 0; i < len(sentence); i++ {
		ch := sentence[i]
		
		index := ch - 'a'
		cheker[index] = true
	}

	for i := 0; i < 26; i++ {
		if(cheker[i] == false) {
			return false
		}
	}
	return true
}