package solution

// Runtime 212 ms  // TODO: 12.5%
// Memory 77.2 MB

func possibleBipartition(N int, dislikes [][]int) bool {
	dis := make([][]int, N+1)
	for i := range dis {
		dis[i] = make([]int, 0, N+1)
	}
	for _, d := range dislikes {
		dis[d[0]] = append(dis[d[0]], d[1])
		dis[d[1]] = append(dis[d[1]], d[0])
	}

	groups := make(map[int]int)
	for i := range dis {
		if _, ok := groups[i]; ok {
			continue
		}
		if !dfs(i, groups, 0, dis) {
			return false
		}
	}
	return true
}

func dfs(i int, groups map[int]int, group int, dis [][]int) bool {
	if g, ok := groups[i]; ok {
		return g == group
	}
	groups[i] = group

	for _, d := range dis[i] {
		if !dfs(d, groups, 1-group, dis) {
			return false
		}
	}
	return true
}
