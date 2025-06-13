package main

import (
	"fmt"
	"math"
)

func toDec(a string) int {
	result := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}

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

func reverse(s string) string {
	sReverse := []rune{}

	for i := len(s) - 1; i >= 0; i-- {
		sReverse = append(sReverse, rune(s[i]))
	}
	return string(sReverse)
}

func addBinary(a string, b string) string {
	aDec := toDec(a)
	bDec := toDec(b)

	cDec := aDec + bDec

	cBin := toBin(cDec)
	return reverse(cBin)
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}