package solution

// Runtime 0 ms
// Memory 2.5 MB

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	x, y := 0, 0
	dx := 0
	dy := 1
	v := 1
	for i := 0; i < n*n; i++ {
		result[x][y] = v
		if x+dy < 0 || x+dx >= n || y+dy < 0 || y+dy >= n || result[x+dx][y+dy] != 0 {
			dx, dy = dy, -dx
		}
		v++
		x += dx
		y += dy
	}

	return result
}
