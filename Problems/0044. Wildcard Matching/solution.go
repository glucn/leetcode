package solution

// Runtime 4 ms, faster than 93.86%
// Memory 3 MB, less than 100%

func isMatch(s string, p string) bool {
	idxS, idxP := 0, 0
	star := -1
	wildMatch := 0

	for idxS < len(s) {
		// the current character in s matches the current character in p
		if idxP < len(p) && (s[idxS] == p[idxP] || p[idxP] == '?') {
			idxS++
			idxP++
			continue
		}

		// the current character in p is the wildcard
		// note down the position of where it is in s, assuming `*` is matching an empty string first,
		// note down the position of the wildcard,
		// and move forward in p
		if idxP < len(p) && p[idxP] == '*' {
			star = idxP
			wildMatch = idxS
			idxP++
			continue
		}

		// the current character in s does not match the current character in p
		// and there was a wildcard in p at `star`
		// move the index in p back to the character after `*`,
		// move the index in s back, and try to match more with `*`
		if star != -1 {
			idxP = star + 1
			wildMatch++
			idxS = wildMatch
			continue
		}

		// there was no wildcard, return false
		return false
	}

	for idxP < len(p) {
		if p[idxP] == '*' {
			idxP++
		} else {
			return false
		}
	}

	return true
}
