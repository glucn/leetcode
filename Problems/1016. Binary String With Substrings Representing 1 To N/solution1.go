package solution

// Runtime 148 ms
// Memory 18.6 MB

// this solution is too slow

func queryString_v1(S string, N int) bool {
	count := 0
	m := make(map[string]struct{})
	for i := 0; i < len(S); i++ {
		if S[i] == '0' {
			continue
		}
		n := 0
		for j := i; j < len(S); j++ {
			s := S[i : j+1]
			n = n<<1 + int(S[j]) - int('0')
			if _, ok := m[s]; !ok {
				m[s] = struct{}{}
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
