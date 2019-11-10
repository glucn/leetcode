package solution

// Runtime 0 ms, faster than ?%
// Memory 2.7 MB, less than ?%

func oddCells(n int, m int, indices [][]int) int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
	}
	for _, idx := range indices {
		for i := 0; i < n; i++ {
			result[i][idx[1]]++
		}
		for j := 0; j < m; j++ {
			result[idx[0]][j]++
		}
	}
	count := 0
	for i := range result {
		for j := range result[i] {
			if result[i][j]%2 == 1 {
				count++
			}
		}
	}
	return count
}
