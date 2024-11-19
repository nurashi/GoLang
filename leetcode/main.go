package main 

import "fmt"


func main() { 
	numbers := []int{2,3,4}

	for _, value := range numbers {
		fmt.Println(value)
	}
}