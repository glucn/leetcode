package solution

// Runtime 56 ms, faster than 46.15%
// Memory 6.8 MB, less than 100%

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	countMap := make(map[string]int)
	result := 0

	letterCount := make(map[byte]int)
	i := 0
	for j := 0; j < len(s); j++ {
		letterCount[s[j]] = letterCount[s[j]] + 1
		if j-i+1 > minSize {
			letterCount[s[i]] = letterCount[s[i]] - 1
			if letterCount[s[i]] == 0 {
				delete(letterCount, s[i])
			}
			i++
		}
		if j-i+1 >= minSize && j-i+1 <= maxSize && len(letterCount) <= maxLetters {
			countMap[s[i:j+1]] = countMap[s[i:j+1]] + 1
			result = max(result, countMap[s[i:j+1]])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
