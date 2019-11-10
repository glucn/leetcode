package _699__Falling_Squares

// Runtime 44 ms, faster than 7.14%
// Memory 7.2 MB, less than 100%

func fallingSquares(positions [][]int) []int {
	var intervals []interval
	var result []int

	maxHeight := 0
	currentHeight := 0
	for _, p := range positions {
		currentHeight, maxHeight = findHeight(intervals, interval{start: p[0], end: p[0] + p[1], height: p[1]}, maxHeight)
		intervals = append(intervals, interval{start: p[0], end: p[0] + p[1], height: currentHeight})
		result = append(result, maxHeight)
	}
	return result
}

type interval struct {
	start  int
	end    int
	height int
}

func findHeight(existing []interval, now interval, maxHeight int) (int, int) {
	bottom := 0
	for _, i := range existing {
		if i.end <= now.start || i.start >= now.end {
			continue
		}
		bottom = max(bottom, i.height)
	}
	nowHeight := bottom + now.height
	return nowHeight, max(maxHeight, nowHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
