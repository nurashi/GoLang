package leedcode


import "strconv"


func addDigits(num int) int {
	if(num / 10 == 0){
		return num
	}

	data := strconv.Itoa(num)
	sum := 0;
	for i := 0; i < len(data); i++{
		sum += int(data[i] - '0')
	}
	return addDigits(sum)
}