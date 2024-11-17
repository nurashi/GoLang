package main

import "fmt"

func main() {
	var n int
	sum := 0;
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i]);
		sum += arr[i]
	}
	avarage := sum / n;
	for i:= 0; i < n; i++ {
		if(arr[i] > avarage){
			fmt.Print(arr[i], " ")
		}
	}	
}
