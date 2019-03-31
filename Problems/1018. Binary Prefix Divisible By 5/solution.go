package solution

// Runtime 336 ms
// Memory 8.2 MB

func prefixesDivBy5(A []int) []bool {
	answer := make([]bool, len(A))
	sum := 0

	for i := range A {
		sum = (sum*2 + A[i]) % 5
		if sum%5 == 0 {
			answer[i] = true
		}
	}
	return answer
}
