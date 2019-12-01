package solution

// Runtime 0 ms, faster than ?%
// Memory 2.5 MB, less than ?%

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 || tomatoSlices < 2*cheeseSlices || tomatoSlices > 4*cheeseSlices {
		return nil
	}
	return []int{(tomatoSlices - 2*cheeseSlices) >> 1, (4*cheeseSlices - tomatoSlices) >> 1}
}
