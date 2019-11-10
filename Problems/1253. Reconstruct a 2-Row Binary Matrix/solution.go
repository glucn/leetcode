package solution

// Runtime 4808 ms, faster than ?%
// Memory 10.7 MB, less than ?%

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	result := make([][]int, 2)
	result[0] = make([]int, len(colsum))
	result[1] = make([]int, len(colsum))

	for i := range colsum {
		if colsum[i] == 2 {
			result[0][i], result[1][i] = 1, 1
			upper--
			lower--
		}
		if upper < 0 || lower < 0 {
			return [][]int{}
		}
	}

	for i := range colsum {
		if colsum[i] == 1 {
			if upper > 0 {
				result[0][i] = 1
				upper--
			} else if lower > 0 {
				result[1][i] = 1
				lower--
			} else {
				return [][]int{}
			}
		}
	}
	if upper != 0 || lower != 0 {
		return [][]int{}
	}
	return result
}
