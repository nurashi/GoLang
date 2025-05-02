package main 

func cheker(word string, letter string) bool {
	count := 0
	for i := 0; i < len(word); i++ {
		if letter == string(word[i]) {
			count++
		}
	}
	return count == 1
}


func firstUniqChar(s string) int {
	
	for i := 0; i < len(s); i++{
		if(cheker(s, string(s[i]))){
			return i;
		}
	}
	return -1;
}

func main() { 
	firstUniqChar("leetcode");

}