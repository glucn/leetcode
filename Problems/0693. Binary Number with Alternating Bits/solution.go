package solution

// Runtime 0 ms
// Memory 2 MB

func hasAlternatingBits(n int) bool {
	last := -1

	for n > 0 {
		digit := n % 2
		if digit == last {
			return false
		}
		last = digit
		n = n >> 1
	}

	return true
}
