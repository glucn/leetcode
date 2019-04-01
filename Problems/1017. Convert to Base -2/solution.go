package solution

// Runtime 0 ms
// Memory 2 MB

func baseNeg2(N int) string {
	var result string
	if N == 0 {
		return "0"
	}
	if N == 1 {
		return "1"
	}
	for N > 1 || N < -1 {
		r := N % (-2)
		if r == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}
		N = -(N >> 1)
	}
	if N == -1 {
		return "11" + result
	} else if N == 1 {
		return "1" + result
	}
	return result
}
