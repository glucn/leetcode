package solution

// Runtime 4 ms, faster than 86.96%
// Memory 5.6 MB, less than 66.67%

func reverseVowels(s string) string {
	var vowelIndexs []int

	for i := range s {
		if isVowel(s[i]) {
			vowelIndexs = append(vowelIndexs, i)
		}
	}

	result := []byte(s)

	for i := 0; i < len(vowelIndexs)>>1; i++ {
		result[vowelIndexs[i]], result[vowelIndexs[len(vowelIndexs)-1-i]] = result[vowelIndexs[len(vowelIndexs)-1-i]], result[vowelIndexs[i]]
	}

	return string(result)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
