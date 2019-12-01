package solution

// Runtime 1264 ms, faster than ?%
// Memory 272.3 MB, less than ?%

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	var result [][]int

	for _, interval := range intervals {
		if interval[0] >= toBeRemoved[1] || interval[1] <= toBeRemoved[0] {
			result = append(result, interval)
			continue
		}
		if interval[0] < toBeRemoved[0] {
			result = append(result, []int{interval[0], toBeRemoved[0]})
		}

		if interval[1] > toBeRemoved[1] {
			result = append(result, []int{toBeRemoved[1], interval[1]})
		}
	}
	return result
}
