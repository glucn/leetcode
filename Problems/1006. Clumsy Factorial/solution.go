package solution

// Runtime 56 ms
// Memory 2.3 MB

func clumsy(N int) int {
	if N == 1 {
		return 1
	}
	if N == 2 {
		return 2
	}
	if N == 3 {
		return 6
	}

	tail := N % 4
	result := 0
	if tail == 1 {
		result = -1
	}
	if tail == 2 {
		result = -2
	}
	if tail == 3 {
		result = -6
	}

	for i := N; i >= 4; i = i - 4 {
		if i == N {
			result = result + i*(i-1)/(i-2) + (i - 3)
		} else {
			result = result - i*(i-1)/(i-2) + (i - 3)
		}
	}
	return result
}
