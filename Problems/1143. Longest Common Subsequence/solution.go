package solution

// Runtime 4 ms
// Memory 10.5 MB

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
	}

	if text1[0] == text2[0] {
		dp[0][0] = 1
	}

	for i := 1; i < len(text1); i++ {
		if dp[i-1][0] == 1 || text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}

	for j := 1; j < len(text2); j++ {
		if dp[0][j-1] == 1 || text1[0] == text2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
