package solution

// Runtime 0 ms, faster than 100%
// Memory 2.2 MB, less than 100%

func longestPalindrome(s string) int {
	countMap := make(map[byte]int)

	for i := range s {
		countMap[s[i]] = countMap[s[i]] + 1
	}

	hasOdd := false
	pairs := 0
	for _, v := range countMap {
		if !hasOdd && v%2 == 1 {
			hasOdd = true
		}
		pairs += v >> 1
	}
	if hasOdd {
		return 2*pairs + 1
	}
	return 2 * pairs
}
