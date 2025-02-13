package leedcode

func countNegatives(grid [][]int) int {

	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				counter++
			}
		}
	}
	return counter
}
