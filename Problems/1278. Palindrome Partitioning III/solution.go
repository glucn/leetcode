package solution

import "math"

// Runtime 24 ms, faster than ?%
// Memory 2.2 MB, less than ?%

func palindromePartition(s string, k int) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := range dp {
		for j := 0; j < k; j++ {
			for next := 1; i+next <= len(s); next++ {
				cost := 0
				for a, b := i, i+next-1; a < b; a, b = a+1, b-1 {
					if s[a] != s[b] {
						cost++
					}
				}

				dp[i+next][j+1] = min(dp[i+next][j+1], dp[i][j]+cost)
			}
		}
	}

	return dp[len(s)][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
