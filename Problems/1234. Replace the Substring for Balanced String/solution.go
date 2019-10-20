package solution

// Runtime 52 ms
// Memory 3.2 MB

func balancedString(s string) int {
	count := make(map[byte]int)
	for i := range s {
		count[s[i]] = count[s[i]] + 1
	}
	// fmt.Println(count)

	l := len(s)
	result := l
	i := 0

	for j := 0; j < len(s); j++ {
		count[s[j]]--
		for i < l && count['Q'] <= l/4 && count['W'] <= l/4 && count['E'] <= l/4 && count['R'] <= l/4 {
			if (j - i + 1) < result {
				result = j - i + 1
			}
			count[s[i]]++
			i++
		}
	}

	return result
}
