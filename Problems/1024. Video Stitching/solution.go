package solution

import "sort"

// Runtime 0 ms
// Memory 2.1 MB

func videoStitching(clips [][]int, T int) int {
	sort.Sort(Clips(clips))
	// fmt.Println(clips)

	count := 0
	end := 0

	for i := range clips {
		j := i
		newEnd := end
		for j < len(clips) && clips[j][0] <= end {
			newEnd = max(newEnd, clips[j][1])
			j++
		}

		if newEnd > end {
			// fmt.Println("new end", newEnd)
			end = newEnd
			count++
			if end >= T {
				return count
			}
		}
	}
	if end >= T {
		return count
	}
	return -1
}

type Clips [][]int

func (c Clips) Len() int {
	return len(c)
}

func (c Clips) Less(i, j int) bool {
	if c[i][0] < c[j][0] {
		return true
	} else if c[i][0] > c[j][0] {
		return false
	} else {
		return c[i][1] < c[j][1]
	}
}

func (c Clips) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
