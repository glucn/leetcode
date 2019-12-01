package solution

// Runtime 280 ms, faster than ?%
// Memory 7 MB, less than ?%

// TODO: this is a brute force solution, find a better solution
func countSquares(matrix [][]int) int {
	maxSide := len(matrix)
	if len(matrix[0]) < len(matrix) {
		maxSide = len(matrix[0])
	}

	count := 0

	for s := 0; s < maxSide; s++ {
		for i := range matrix {
			for j := range matrix[i] {
				if i+s >= len(matrix) || j+s >= len(matrix[0]) {
					continue
				}
				allOne := true
			loop:
				for x := i; x <= i+s; x++ {
					for y := j; y <= j+s; y++ {
						if matrix[x][y] == 0 {
							allOne = false
							break loop
						}
					}
				}
				if allOne {
					count++
				}
			}
		}
	}

	return count

}
