package main

import "fmt"

func main() { 
	counter := 0
	var arr[5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
		if arr[i] > 0 {
			counter++
		}
	}
	fmt.Print(counter)
}