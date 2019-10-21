package solution

// Runtime 12 ms
// Memory 6.4 MB

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, 0

	for j < len(t) {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}

		j++
	}
	return false
}
