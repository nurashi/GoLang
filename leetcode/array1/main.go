package main

import "fmt"

func isValid(s string) bool {

	helper := false

	if len(s) <= 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i]+2 == s[j] || s[i]+1 == s[j] {
				helper = true
			}
		}
	}

	return helper
}

func main() {
	s := "()"

	fmt.Println(isValid(s))
}
