package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num < 1 || num > 12 {
		fmt.Println("Unreal month")
	} else if num < 3 || num == 12 {
		fmt.Println("Winter")
	} else if num >= 3 && num <= 5 {
		fmt.Println("Spring")
	} else if num >= 6 && num <= 8 {
		fmt.Println("Summer")
	} else {
		fmt.Println("Autumn")
	}
}
