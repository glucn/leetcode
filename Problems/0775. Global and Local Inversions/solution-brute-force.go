package solution

// Runtime 432 ms
// Memory 6.3 MB

func isIdealPermutationBruteForce(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	countLocal := countLocal(A)
	countGlobal := countGlobal(A)
	return countLocal == countGlobal
}

func countLocal(A []int) int {
	count := 0
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			count++
		}
	}
	return count
}

func countGlobal(A []int) int {
	count := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				count++
			}
		}
	}
	return count
}
