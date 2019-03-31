package solution

// Runtime 256 ms
// Memory 24.8 MB

import "sort"

//Definition for an interval.
type Interval struct {
	Start int
	End   int
}

type intervalWithIndex struct {
	interval Interval
	index    int
}

type byLeftPoint []intervalWithIndex

func (b byLeftPoint) Len() int {
	return len(b)
}

func (b byLeftPoint) Less(i, j int) bool {
	return b[i].interval.Start < b[j].interval.Start
}

func (b byLeftPoint) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func findRightInterval(intervals []Interval) []int {
	length := len(intervals)
	withIndex := make([]intervalWithIndex, length)

	for i, inter := range intervals {
		withIndex[i] = intervalWithIndex{
			interval: inter,
			index:    i,
		}
	}

	sort.Sort(byLeftPoint(withIndex))

	output := make([]int, length)

	for i := 0; i < length; i++ {
		output[withIndex[i].index] = bSearch(withIndex, withIndex[i].interval.End)
	}
	return output
}

func bSearch(intervals []intervalWithIndex, target int) int {
	if target > intervals[len(intervals)-1].interval.Start {
		return -1
	}

	start, end := 0, len(intervals)-1
	result := len(intervals) - 1

	for start <= end {
		middle := (start + end) / 2

		if target == intervals[middle].interval.Start {
			result = middle
			break
		} else if target < intervals[middle].interval.Start {
			result = middle
			end = middle - 1
		} else {
			start = middle + 1
		}

	}
	return intervals[result].index
}
