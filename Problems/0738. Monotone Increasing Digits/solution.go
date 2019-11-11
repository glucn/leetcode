package solution

import "strconv"

// Runtime 0 ms, faster than 100%
// Memory 2 MB, less than 100%

func monotoneIncreasingDigits(N int) int {
	b := []byte(strconv.Itoa(N))
	i := 1
	for i < len(b) && b[i] >= b[i-1] {
		i++
	}

	for i > 0 && i < len(b) && b[i] < b[i-1] {
		b[i-1]--
		i--
	}

	for j := i + 1; j < len(b); j++ {
		b[j] = '9'
	}

	result, _ := strconv.Atoi(string(b))
	return result
}
