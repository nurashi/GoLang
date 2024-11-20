package main 

import "fmt"


func main() { 
	numbers := []int{2,3,4}

	for index, value := range numbers {
		fmt.Println(index, " ", value)
	}
}