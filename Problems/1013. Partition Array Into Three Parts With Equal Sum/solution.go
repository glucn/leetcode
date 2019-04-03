package solution

// Runtime 48 ms
// Memory 6.8 MB

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}

	total := 0
	sum := make([]int, len(A))
	for i := range A {
		total += A[i]
		if i == 0 {
			sum[i] = A[0]
		} else {
			sum[i] = sum[i-1] + A[i]
		}
	}

	if total%3 != 0 {
		return false
	}
	target := total / 3

	left, right := 0, len(A)-1
	for left < right {
		if sum[left] != target {
			left++
			continue
		}
		if sum[right] != 2*target {
			right--
			continue
		}
		return true
	}
	return false
}
