package solution

// Runtime 20 ms, faster than 88.14%
// Memory 7.3 MB, less than 100%

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left+1:right+1]) || isPalindrome(s[left:right])
		}
		left, right = left+1, right-1
	}
	return true
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}
