package solution

// Runtime 4 ms, faster than 91.28%
// Memory 3.1 MB, less than 100%

func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		if isEven(n) {
			count++
		}
	}
	return count
}

func isEven(num int) bool {
	if num < 10 {
		return false
	}
	return !isEven(num / 10)
}
