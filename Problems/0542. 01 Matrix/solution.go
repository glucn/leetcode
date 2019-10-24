package solution

// Runtime 228 ms
// Memory 93.1 MB

import "math"

func updateMatrix(matrix [][]int) [][]int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		for j := range dp[i] {
			if matrix[i][j] != 0 {
				dp[i][j] = math.MaxInt32
			}
		}
	}

	for i := 0; i<len(matrix); i++ {
		for j := 0; j<len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				if i > 0 {
					dp[i][j] = min(dp[i-1][j]+1, dp[i][j])
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j-1]+1, dp[i][j])
				}
			}
		}
	}

	for i := len(matrix)-1; i>=0; i-- {
		for j := len(matrix[0])-1; j>=0; j-- {
			if matrix[i][j] != 0 {
				if i < len(matrix)-1 {
					dp[i][j] = min(dp[i+1][j]+1, dp[i][j])
				}
				if j < len(matrix[0])-1 {
					dp[i][j] = min(dp[i][j+1]+1, dp[i][j])
				}
			}
		}
	}

	return dp
}

func min(nums ...int) int {
	min := math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
