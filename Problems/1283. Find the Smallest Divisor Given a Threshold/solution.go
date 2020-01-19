package solution

// Runtime 44 ms, faster than 40.74%
// Memory 7.3 MB, less than 100%

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, int(1e6)
	for left < right {
		mid := (left + right) >> 1
		sum := 0
		for _, n := range nums {
			sum += divide(n, mid)
		}
		if sum > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func divide(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}
