package solution

// Runtime 1068 ms
// Memory 8.8 MB

func lexicalOrder(n int) []int {
	if n < 1 {
		return []int{}
	}

	var result []int
	for i := 1; i < 10; i++ {
		result = append(result, dfs(i, n)...)
	}
	return result
}

func dfs(leading int, n int) (result []int) {
	if leading > n {
		return
	}
	result = []int{leading}
	for i := 0; i < 10; i++ {
		next := 10*leading + i
		if next > n {
			return
		}
		result = append(result, dfs(next, n)...)
	}
	return
}
