package solution

import "strconv"
import "sort"
import "strings"

// Runtime 4ms
// Memory 2.6 MB

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	var output string
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(customSort(strs))
	if strs[0] == "0" {
		return "0"
	}
	for i := range strs {
		output = output + strs[i]
	}
	return output
}

type customSort []string

func (c customSort) Len() int {
	return len(c)
}

func (c customSort) Less(i, j int) bool {
	return comp(c[i], c[j])
}

func comp(x, y string) bool {
	if len(x) == 0 {
		return false
	}
	if len(y) == 0 {
		return true
	}
	if strings.Index(x, y) == 0 {
		return comp(x[len(y):], y)
	}

	if strings.Index(y, x) == 0 {
		return comp(x, y[len(x):])
	}
	return x > y
}

func (c customSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
