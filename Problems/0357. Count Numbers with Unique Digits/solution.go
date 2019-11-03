package solution

// Runtime 0 ms
// Memory 1.9 MB

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	if n > 10 {
		return countNumbersWithUniqueDigits(10)
	}
	total := 10
	current := 9 // the first digit cannot be 0

	for i := 9; i >= 9-n+2; i-- {
		current *= i
		total += current
	}
	return total
}
