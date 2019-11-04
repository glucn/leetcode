package solution

// Runtime 4 ms, faster than 97.37%
// Memory 3.7 MB, less than 100%

// a case which should not be ignored: [[1,0],[0,2]]

func projectionArea(grid [][]int) int {
	maxX := make([]int, len(grid))
	maxY := make([]int, len(grid[0]))
	countZero := 0

	for i := range grid {
		for j := range grid[i] {
			maxX[i] = max(maxX[i], grid[i][j])
			maxY[j] = max(maxY[j], grid[i][j])
			if grid[i][j] == 0 {
				countZero++
			}
		}
	}

	sum := len(grid)*len(grid[0]) - countZero
	for i := range maxX {
		sum += maxX[i]
	}
	for i := range maxY {
		sum += maxY[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
