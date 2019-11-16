package solution

// Runtime 4 ms, faster than 95.61%
// Memory 44.3 MB, less than 100%

func canConstruct(ransomNote string, magazine string) bool {
	count := make([]int, 26)
	for i := range magazine {
		count[magazine[i]-'a']++
	}
	for i := range ransomNote {
		count[ransomNote[i]-'a']--
		if count[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
