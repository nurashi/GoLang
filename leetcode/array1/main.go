package main

import "fmt"

func removeDuplicates(gcdStr []rune) string {

	seen := make(map[rune]bool)
	var result []rune
	for _, i := range gcdStr {
		if !seen[i] {
			seen[i] = true
			result = append(result, i)
		}
	}
	return string(result)
}

func gcdOfStrings(str1 string, str2 string) string {

	gcdStr := []rune{}
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				gcdStr = append(gcdStr, rune(str1[i]))
			}
		}
	}
	result := removeDuplicates(gcdStr)
	return result
}

func main() {
	str1 := "LEET"
	str2 := "CODE"
	fmt.Println(gcdOfStrings(str1, str2))
}
