package solution

// Runtime 0 ms, faster than 100%
// Memory 1.9 MB, less than 199%

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}

	var result []byte
	for n > 0 {
		d := byte((n-1)%26) + 'A'
		result = append(result, d)
		n = (n - 1) / 26
	}
	return string(reverse(result))
}

func reverse(s []byte) []byte {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
	return s
}
