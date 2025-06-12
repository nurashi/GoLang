package leedcode


func numberOfEmployeesWhoMetTarget(hours []int, target int) int {

	counter := 0
	for i := 0; i <len(hours); i++ { 
		if(hours[i] >= target) {
			counter++
		}
	}
	return counter
}