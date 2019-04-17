package solution

// Runtime 952 ms
// Memory 102.8 MB

// TODO: can be faster

func longestArithSeqLength(A []int) int {
	m := make(map[int]map[int]int) // diff -> index -> count
	res := 0

	for i := range A {
		for j := i + 1; j < len(A); j++ {
			a, b := A[i], A[j]
			if _, ok := m[b-a]; !ok {
				// fmt.Println("new map", i, b-a)
				m[b-a] = make(map[int]int)
				m[b-a][j] = 1
			} else {
				// fmt.Println("increment map", j, b-a)
				m[b-a][j] = max(m[b-a][j], m[b-a][i]+1)
				res = max(res, m[b-a][j]+1)
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
