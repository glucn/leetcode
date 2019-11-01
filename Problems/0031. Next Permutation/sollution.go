package solution

// Runtime 0 ms
// Memory 3.4 MB

import "sort"

func nextPermutation(nums []int) {
	dropIdx := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			dropIdx = i - 1
			break
		}
	}

	if dropIdx == -1 {
		sort.Sort(sort.IntSlice(nums))
		return
	}

	for i := len(nums) - 1; i > dropIdx; i-- {
		if nums[i] > nums[dropIdx] {
			nums[i], nums[dropIdx] = nums[dropIdx], nums[i]
			sort.Sort(sort.IntSlice(nums[dropIdx+1:]))
			return
		}
	}

}
