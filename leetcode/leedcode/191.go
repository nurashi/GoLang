package leedcode


func toBin(dec int) string {
	if dec == 0 {
		return "0"
	}
	var runes []rune
	for dec > 0 {
		if dec % 2 == 0 {
			runes = append(runes, '0')
		} else {
			runes = append(runes, '1')
		}
		dec /= 2
	}

	return string(runes)
}

func count(s string) int {
	counter := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			counter++
		}
	}
	return counter
}

func hammingWeight(n int) int {
	return count(toBin(n))
}