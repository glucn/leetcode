package solution

// Runtime 0 ms
// Memory 2 MB

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}
	fives := n / 5

	return fives + trailingZeroes(n/5)
}
