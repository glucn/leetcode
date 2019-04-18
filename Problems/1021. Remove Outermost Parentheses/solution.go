package solution

// Runtime 0 ms
// Memory 2.5 MB

func removeOuterParentheses(S string) string {
	result := make([]byte, len(S))
	layer := 0
	count := 0

	for i:= 0; i<len(S); i++ {
		if S[i] == '(' {
			if layer != 0 {
				result[count] = '('
				count++
			}
			layer++
		} else {
			layer--
			if layer != 0 {
				result[count] = ')'
				count++
			}
		}
	}

	return string(result[:count])
}
