package solution

import "math"

// Runtime 16 ms
// Memory 2.7 MB

func largestSumAfterKNegations(A []int, K int) int {
	output := make([]int, K)
	for i := 0; i < K; i++ {
		output[i] = findIndex(A)
		A[output[i]] = -A[output[i]]
	}

	sum := 0
	for _, a := range A {
		sum += a
	}
	return sum
}

func findIndex(A []int) int {
	minPositiveIndex, minNagativeIndex := -1, -1
	minPositive, minNegative := math.MaxInt32, 0

	for i, a := range A {
		if a < minNegative {
			minNegative = a
			minNagativeIndex = i
		}
		if a >= 0 && a < minPositive {
			minPositive = a
			minPositiveIndex = i
		}
	}

	if minNagativeIndex >= 0 {
		return minNagativeIndex
	}
	return minPositiveIndex
}
