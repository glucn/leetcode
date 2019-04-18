package solution

import "strings"

// Runtime 0 ms
// Memory 2.1 MB

func queryString_v2(S string, N int) bool {
	if N == 0 {
		return true
	}
	s := iToB(N)
	return strings.Index(S, s) > -1 && queryString_v2(S, N-1)
}

func iToB(n int) string {
	var res []byte
	for n != 0 {
		if n%2 == 0 {
			res = append([]byte{'0'}, res...)
		} else {
			res = append([]byte{'1'}, res...)
		}
		n = n >> 1
	}
	return string(res)
}
