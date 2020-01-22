package solution

import "strings"

// Runtime 196 ms, faster than 88.89%
// Memory 7.2 MB, less than 100%

func maxEqualRowsAfterFlips(matrix [][]int) int {
	count := make(map[string]int)
	for i := range matrix {
		origin := strings.Builder{}
		flip := strings.Builder{}
		for j := range matrix[i] {
			origin.WriteByte(byte(matrix[i][j] + '0'))
			flip.WriteByte(byte(1 - matrix[i][j] + '0'))
		}
		count[origin.String()] = count[origin.String()] + 1
		if origin.String() != flip.String() {
			count[flip.String()] = count[flip.String()] + 1
		}
	}

	max := 0
	for _, c := range count {
		if c > max {
			max = c
		}
	}
	return max
}
