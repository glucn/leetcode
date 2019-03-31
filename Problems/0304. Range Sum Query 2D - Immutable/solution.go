package solution

// Runtime 44 ms
// Memory 9.5 MB

type NumMatrix struct {
	matrix   [][]int
	cacheSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	cache := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		cache[i] = make([]int, len(matrix[0])+1)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			cache[i+1][j+1] = cache[i+1][j] + cache[i][j+1] + matrix[i][j] - cache[i][j]
		}
	}

	return NumMatrix{
		matrix:   matrix,
		cacheSum: cache,
	}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return m.cacheSum[row2+1][col2+1] - m.cacheSum[row1][col2+1] - m.cacheSum[row2+1][col1] + m.cacheSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
