package solution

// Runtime 12 ms
// Memory 6.8 MB

func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0, len(s))
	toRemove := make(map[int]struct{})
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		}
		if s[i] == ')' {
			if len(stack) == 0 {
				toRemove[i] = struct{}{}
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for _, s := range stack {
		toRemove[s] = struct{}{}
	}

	// fmt.Println(toRemove)

	result := make([]byte, 0, len(s))
	for i := range s {
		if _, ok := toRemove[i]; !ok {
			result = append(result, s[i])
		}
	}
	return string(result)
}
