package solution

// Runtime 72 ms
// Memory 6.3 MB

func isIdealPermutation(A []int) bool {
	for i := range A {
		if abs(A[i]-i) > 1 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
