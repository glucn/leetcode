package solution

// Runtime 0 ms, faster than 100%
// Memory 2.1 MB, less than 100%

func minCostToMoveChips(chips []int) int {
	countOdd, countEven := 0, 0

	for _, c := range chips {
		if c%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}

	return min(countOdd, countEven)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
