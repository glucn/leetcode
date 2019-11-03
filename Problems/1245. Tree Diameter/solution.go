package solution

// Runtime 28 ms
// Memory 7.8 MB

func treeDiameter(edges [][]int) int {
	if len(edges) == 0 {
		return 0
	}
	m := make(map[int][]int)
	for _, e := range edges {
		m[e[0]] = append(m[e[0]], e[1])
		m[e[1]] = append(m[e[1]], e[0])
	}

	// start with a random node, find the node which is the farthest from the starting node.
	// it must be one end of the longest path
	start, _ := bfs(0, m)

	// find another end
	_, dis := bfs(start, m)

	return dis
}

func bfs(start int, m map[int][]int) (farthestNode int, farthestDistance int) {
	queue := []int{start}
	distances := make([]int, len(m))
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
			visited[nn] = true
		}
	}

	for i, d := range distances {
		if d > farthestDistance {
			farthestDistance = d
			farthestNode = i
		}
	}

	return farthestNode, farthestDistance
}
