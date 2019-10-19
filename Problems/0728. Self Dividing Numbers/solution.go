package solution

// Runtime 0 ms
// Memory 2.7 MB

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for num := left; num <= right; num++ {
		if canSelfDividing(num) {
			result = append(result, num)
		}
	}
	return result
}

func canSelfDividing(num int) bool {
	n := num
	for n > 0 {
		d := n % 10
		if d == 0 || num%d != 0 {
			return false
		}
		n = n / 10
	}
	return true
}
