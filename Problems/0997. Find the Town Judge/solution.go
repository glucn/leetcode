package solution

// Runtime 100 ms
// Memory 7.4 MB

func findJudge(N int, trust [][]int) int {
	count := make([]int, N)
	for _, t := range trust {
		if count[t[1]-1] != -1 {
			count[t[1]-1]++
		}
		count[t[0]-1] = -1
	}

	for i := range count {
		if count[i] == N-1 {
			return i+1
		}
	}
	return -1
}
