package solution

import "sort"

// Runtime 64 ms, faster than ?%
// Memory 7 MB, less than ?%

func maxSumDivThree(nums []int) int {
	total := 0
	totalThree := 0
	one := []int{}
	two := []int{}

	for _, n := range nums {
		total += n
		if n%3 == 1 {
			one = append(one, n)
		} else if n%3 == 2 {
			two = append(two, n)
		} else {
			totalThree += n
		}
	}

	if total%3 == 0 {
		return total
	}

	sort.Sort(sort.IntSlice(one))
	sort.Sort(sort.IntSlice(two))
	if total%3 == 1 {
		if len(one) == 0 && len(two) < 2 {
			return totalThree
		}
		if len(one) == 0 {
			return total - two[0] - two[1]
		}
		if len(two) < 2 {
			return total - one[0]
		}
		return total - min(one[0], two[0]+two[1])
	}

	// total % 3 == 2
	if len(two) == 0 && len(one) < 2 {
		return totalThree
	}
	if len(two) == 0 {
		return total - one[0] - one[1]
	}
	if len(one) < 2 {
		return total - two[0]
	}
	return total - min(two[0], one[0]+one[1])

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
