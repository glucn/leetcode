package solution

import "sort"

// Runtime 208 ms
// Memory 219.3 MB

func minimumAbsDifference(arr []int) [][]int {
	if len(arr) <= 1 {
		return nil
	}
	sort.Sort(sort.IntSlice(arr))

	minDiff := 2000001 // 2*10^6 +1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < minDiff {
			minDiff = arr[i] - arr[i-1]
		}
	}

	var result [][]int
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minDiff {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
