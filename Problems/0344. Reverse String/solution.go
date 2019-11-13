package solution

// Runtime 656 ms, faster than 64.73%
// Memory 8.3 MB, less than 100%

func reverseString(s []byte) {
	for i := 0; i < len(s)>>1; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return
}
