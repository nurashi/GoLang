package leedcode



// func toBin(dec int) string {
// 	if dec == 0 {
// 		return "0"
// 	}
// 	var runes []rune
// 	for dec > 0 {
// 		if dec % 2 == 0 {
// 			runes = append(runes, '0')
// 		} else {
// 			runes = append(runes, '1')
// 		}
// 		dec /= 2
// 	}

// 	return string(runes)
// }

func count1(bin string) int {
	counter := 0

	for i := 0; i < len(bin); i++ {
		if(bin[i] == '1'){
			counter++
		}
	}
	return counter
}


func countBits(n int) []int {

	result := []int{}
	for i := 0; i <= n; i++ {
		binNum := toBin(i)
		counted := count1(binNum)

		result = append(result, counted)
	}


	return result
}