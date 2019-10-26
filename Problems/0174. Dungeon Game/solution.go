package solution

import "math"

// Runtime 4 ms
// Memory 3.7 MB

func calculateMinimumHP(dungeon [][]int) int {
	dp := make([][]int, len(dungeon))
	for i := range dp {
		dp[i] = make([]int, len(dungeon[0]))
		for j := range dp[i] {
			dp[i][j] = -math.MaxInt32
		}
	}
	dp[len(dungeon)-1][len(dungeon[0])-1] = min(dungeon[len(dungeon)-1][len(dungeon[0])-1], 0)

	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			if i < len(dungeon)-1 {
				dp[i][j] = min(max(dp[i][j], dp[i+1][j]+dungeon[i][j]), 0)
			}
			if j < len(dungeon[0])-1 {
				dp[i][j] = min(max(dp[i][j], dp[i][j+1]+dungeon[i][j]), 0)
			}
		}
		// fmt.Println(dp)
	}

	return -dp[0][0] + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
