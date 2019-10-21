package solution

import "sort"

// Runtime 76 ms
// Memory 6.7 MB

func arrayPairSum(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
