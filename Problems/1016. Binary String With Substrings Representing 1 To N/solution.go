package solution

// Runtime 288 ms
// Memory 17.4 MB

// TODO: faster than 5.66%....too slow

func queryString(S string, N int) bool {
	count := 0
	m := make(map[string]struct{})
	for i := 0; i < len(S); i++ {
		if S[i] == '0' {
			continue
		}
		for j := i; j < len(S); j++ {
			s := S[i : j+1]
			if _, ok := m[s]; !ok {
				m[s] = struct{}{}
				n := bToI(s)
				if n >= 1 && n <= N {
					count++
					if count == N {
						return true
					}
				}
			}
		}
	}
	return false
}

func bToI(s string) int {
	n := 0
	for i := range s {
		n = n*2 + int(s[i]) - int('0')
	}
	return n
}
