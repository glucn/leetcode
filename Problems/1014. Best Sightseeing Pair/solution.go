package solution

// Runtime 64 ms
// Memory 7.9 MB

// TODO: there is a better solution

func maxScoreSightseeingPair(A []int) int {
	max := 0

	start := make([]int, len(A))
	for i := range A {
		if i == 0 {
			start[i] = A[i]
		} else {
			if A[i]+i > start[i-1] {
				start[i] = A[i] + i
			} else {
				start[i] = start[i-1]
			}
		}
	}
	for j := len(A) - 1; j > 0; j-- {
		if A[j]-j+start[j-1] > max {
			max = A[j] - j + start[j-1]
		}
	}
	return max
}
