package solution

import "math/bits"

// Runtime 44 ms, faster than 97.3%
// Memory 2 MB, less than 100%

func countPrimeSetBits(L int, R int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19} // R <= 10^6
	count := 0

	for i := L; i <= R; i++ {
		if isInSlice(bits.OnesCount(uint(i)), primes) {
			count++
		}
	}

	return count
}

func isInSlice(x int, s []int) bool {
	for _, ss := range s {
		if ss == x {
			return true
		}
	}
	return false
}
