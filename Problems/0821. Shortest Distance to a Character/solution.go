package solution

// Runtime 4ms
// Memory 6.6 MB

// TODO: there is a better solution with O(N) complexity

func shortestToChar(S string, C byte) []int {
	last := -1
	result := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			result[i] = 0
			for j := last + 1; j < i; j++ {
				if i-j < result[j] || last == -1 {
					result[j] = i - j
				}
			}
			last = i
		} else {
			result[i] = i - last
		}
	}
	return result
}
