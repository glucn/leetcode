package solution

// Runtime 20 ms, faster than 87.76%
// Memory 6.2 MB, less than 100%

func findCircleNum(M [][]int) int {
	graph := make(map[int][]int)

	for i := range M {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	visited := make([]bool, len(M))
	count := 0
	for i := range M {
		if visited[i] {
			continue
		}
		count++
		dfs(graph, visited, i)
		visited[i] = true
	}
	return count
}

func dfs(graph map[int][]int, visited []bool, i int) {
	if visited[i] {
		return
	}
	visited[i] = true
	for _, n := range graph[i] {
		dfs(graph, visited, n)
	}
}
