package solution

import "math"

// Runtime 12 ms, faster than 50%
// Memory 2.7 MB, less than 100%

func getMoneyAmount(n int) int {
	if n <= 1 {
		return 0
	}

	dpResult := make([][]int, n+1)
	for i := range dpResult {
		dpResult[i] = make([]int, n+1)
	}

	return dp(1, n, dpResult)
}

func dp(from, to int, dpResult [][]int) int {
	if from >= to {
		return 0
	}
	if dpResult[from][to] != 0 {
		return dpResult[from][to]
	}

	minCost := math.MaxInt32

	for x := from; x <= to; x++ {
		// when I pick x and it's wrong, the answer could be in [from, x-1] or [x+1, to]
		cost := x + max(dp(from, x-1, dpResult), dp(x+1, to, dpResult))
		minCost = min(minCost, cost)
	}

	dpResult[from][to] = minCost
	return minCost

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
