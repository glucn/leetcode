package _699__Falling_Squares

// Runtime 32 ms, faster than 35.71%
// Memory 7.2 MB, less than 100%

func fallingSquares(positions [][]int) []int {
	var intervals []interval
	var result []int

	for _, p := range positions {
		currentHeight := findHeight(intervals, interval{start: p[0], end: p[0] + p[1], height: p[1]})
		intervals = append(intervals, interval{start: p[0], end: p[0] + p[1], height: currentHeight})

		if len(result) > 0 {
			result = append(result, max(result[len(result)-1], currentHeight))
		} else {
			result = []int{currentHeight}
		}
	}
	return result
}

type interval struct {
	start  int
	end    int
	height int
}

func findHeight(existing []interval, now interval) int {
	bottom := 0
	for _, i := range existing {
		if i.end <= now.start || i.start >= now.end {
			continue
		}
		bottom = max(bottom, i.height)
	}
	return bottom + now.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
