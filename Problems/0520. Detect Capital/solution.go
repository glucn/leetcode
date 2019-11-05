package solution

// Runtime 0 ms, faster than 100%
// Memory 2.1 MB, less than 100%

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}
	firstCap := isCap(word[0])
	allCap := isCap(word[1])
	if !firstCap && allCap {
		return false
	}
	for i := 1; i < len(word); i++ {
		if allCap && word[i] >= byte('a') && word[i] <= byte('z') {
			return false
		}
		if !allCap && isCap(word[i]) {
			return false
		}
	}
	return true
}

func isCap(c byte) bool {
	return c >= byte('A') && c <= byte('Z')
}
