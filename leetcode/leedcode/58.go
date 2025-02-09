package leedcode

import "strings"

func lastWordLen(s string) int {
	s = strings.TrimRight(s, " ")

	if len(s) == 1 {
		return len(s)
	}
	counterspace := 0
	counter := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			counterspace++
			counter = i
			break
		}
	}
	if counterspace > 0 {
		return len(s) - counter - 1
	} else {
		return len(s) - counter
	}
}
