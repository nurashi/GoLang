package main

import (
	"fmt"
	"strconv"
)

func countDigits(num int) int {
    strNum := strconv.Itoa(num)
	counter := 0
	for i := 0; i < len(strNum); i++ {
		current := int(strNum[i])
		if(num % current == 0){
			counter++
		}
	}
	return counter + 1;
}

func main() { 
	num := 121
	fmt.Println(countDigits(num))
}