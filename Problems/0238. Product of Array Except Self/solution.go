package solution

// Runtime 2136 ms // TODO: too slow!
// Memory 8.6 MB

func productExceptSelf(nums []int) []int {
	results := make([]int, len(nums))
	results[0] = 1
	for i := 1; i < len(nums); i++ {
		results[i] = results[i-1] * nums[i-1]
	}

	multiple := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		results[i] = results[i] * multiple
		multiple = multiple * nums[i]
	}

	return results
}
