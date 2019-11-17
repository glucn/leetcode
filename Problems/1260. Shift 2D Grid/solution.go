package solution

// Runtime 48 ms, faster than ?%
// Memory 6.8 MB, less than ?%

// TODO: this is not the best solution
func shiftGrid(grid [][]int, k int) [][]int {
	row := len(grid)
	col := len(grid[0])

	if k%(row*col) == 0 {
		return grid
	}

	result := make([][]int, row)
	for i := range result {
		result[i] = make([]int, col)
	}

	for i := range grid {
		result[i][0] = grid[(i-1+row)%row][col-1]
	}
	for i := range grid {
		for j := 1; j < col; j++ {
			result[i][j] = grid[i][j-1]
		}
	}
	return shiftGrid(result, (k-1)%(row*col))
}
