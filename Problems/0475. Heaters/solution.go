package solution

import "math"

// Runtime 2244 ms, faster than 5.77%
// Memory 7.2 MB, less than 100%
// TODO: make it faster

func findRadius(houses []int, heaters []int) int {
	ranges := make([]int, len(houses))
	for i := range ranges {
		ranges[i] = math.MaxInt32
	}

	for i := range houses {
		for j := range heaters {
			ranges[i] = min(ranges[i], abs(houses[i]-heaters[j]))
		}
	}

	result := 0
	for i := range ranges {
		result = max(result, ranges[i])
	}
	return result
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
