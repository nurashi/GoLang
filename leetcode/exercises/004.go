package main

import "fmt"

func main() { 
	var arr[5]int

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
		if arr[i] % 2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
}