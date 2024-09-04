package main

import "fmt"

func main() {
	for i := 2; i < 100; i++ {
		cheker := true
		for j := 2; j*j < i; j++ {
			if i%j == 0 {
				cheker = false
				break
			}
		}
		if cheker {
			fmt.Print(i, " ")
		}
	}
}
