package leedcode

func diagonalSum(mat [][]int) int {
	sum := 0
	size := len(mat)

	if size == 0 {
		return 0
	}

	for i := 0; i < size; i++ {
		sum += mat[i][i]
		if i != size-i-1 {
			sum += mat[i][size-i-1]
		}
	}

	return sum
}
