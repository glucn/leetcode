package solution

// Runtime 336 ms, faster than ?%
// Memory 6.2 MB, less than ?%

func deleteTreeNodes(nodes int, parent []int, value []int) int {
	sum := make([]int, nodes)
	count := make([]int, nodes)

	sum[0] = value[0]
	count[0] = 1

	for i := 1; i < nodes; i++ {
		p := i
		for p != -1 {
			sum[p] += value[i]
			count[p] += 1
			p = parent[p]
		}
	}

	for i := nodes - 1; i > 0; i-- {
		if sum[i] == 0 {
			p := parent[i]
			for p != -1 {
				count[p] -= count[i]
				p = parent[p]
			}
		}
	}

	return count[0]
}
