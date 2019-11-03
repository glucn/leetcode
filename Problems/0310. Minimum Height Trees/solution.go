package solution

// Runtime 0 ms
// Memory 8 MB

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	m := make(map[int][]int, n)
	for _, e := range edges {
		m[e[0]] = append(m[e[0]], e[1])
		m[e[1]] = append(m[e[1]], e[0])
	}

	// start with a random node, find the node which is the farthest from the starting node.
	// it must be one end of the longest path
	start, _, _ := bfs(0, m)

	// find another end
	end, dis, pre := bfs(start, m)

	// form the longest path
	var longestPath []int
	for end != -1 {
		longestPath = append(longestPath, end)
		end = pre[end]
	}

	if dis%2 == 0 {
		return []int{longestPath[(dis >> 1)]}
	}

	return []int{longestPath[(dis >> 1)], longestPath[(dis>>1)+1]}
}

func bfs(start int, m map[int][]int) (farthestNode int, farthestDistance int, previousNode []int) {
	queue := []int{start}
	distances := make([]int, len(m))
	previousNode = make([]int, len(m))
	previousNode[start] = -1
	visited := make([]bool, len(m))
	visited[start] = true

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, nn := range m[n] {
			if visited[nn] {
				continue
			}
			queue = append(queue, nn)
			distances[nn] = distances[n] + 1
			previousNode[nn] = n
			visited[nn] = true
		}
	}

	for i, d := range distances {
		if d > farthestDistance {
			farthestDistance = d
			farthestNode = i
		}
	}

	return farthestNode, farthestDistance, previousNode
}
