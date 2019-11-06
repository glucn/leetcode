package solution

import "math"

// Runtime 296 ms, faster than 15.441% // TODO: improve this
// Memory 4.7 MB, less than 66.67%


func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := range nums {
		for j := 1; j <= nums[i]; j++ {
			if i+j == len(nums) {
				break
			}
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}

	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
