package solution

// Runtime 48 ms, faster than 19.2%
// Memory 5.7 MB, less than 75%
// TODO: improve the solution

func firstUniqChar(s string) int {
	m := make(map[byte]int)

	for k := byte('a'); k <= byte('z'); k++ {
		m[k] = -1
	}

	for i := range s {
		if idx := m[s[i]]; idx != -1 {
			m[s[i]] = len(s) + 1 // duplicate
		} else {
			m[s[i]] = i
		}
	}

	min := len(s)
	for _, v := range m {
		if v < min && v != -1 {
			min = v
		}
	}
	if min == len(s) {
		return -1
	}
	return min
}
