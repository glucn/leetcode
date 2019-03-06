package solution

// Runtime 28 ms
// Memory 6.2 MB

func repeatedNTimes(A []int) int {
	numbers := make(map[int]struct{})

	for _, a := range A {
		if _, found := numbers[a]; found {
			return a
		}
		numbers[a] = struct{}{}
	}
	return -1
}
