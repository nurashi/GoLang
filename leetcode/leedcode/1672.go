package leedcode


func maximumWealth(accounts [][]int) int {
	max := 0
	checkMax := 0
	for i := 0; i < len(accounts); i++ {
		max = 0
		for j := 0; j < len(accounts[i]); j++ {
			max += accounts[i][j]
		}
		if max > checkMax {
			checkMax = max
		}
	}

	return checkMax
}
