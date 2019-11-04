package solution

// Runtime 56 ms, faster than 10.72%
// Memory 5.7 MB, less than 75%
// TODO: improve the solution

func firstUniqChar_v2(s string) int {
	m := make(map[byte]int)

	for i := range s {
		m[s[i]] = m[s[i]] + 1
	}

	for i := range s {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}
