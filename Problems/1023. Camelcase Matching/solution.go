package solution

// Runtime 0 ms
// Memory 2 MB

func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = match(q, pattern)
	}
	return result
}

func match(s string, pattern string) bool {
	ptIdx := 0
	result := false

	for i := range s {
		if ptIdx < len(pattern) && s[i] == pattern[ptIdx] {
			ptIdx++
			if ptIdx == len(pattern) {
				result = true
			}
			continue
		}
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return result
}
