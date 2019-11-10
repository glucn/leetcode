package _255__Maximum_Score_Words_Formed_by_Letters

// Runtime 4 ms, faster than ?%
// Memory 2.4 MB, less than ?%

func maxScoreWords(words []string, letters []byte, score []int) int {
	scores := make([]int, len(words))
	counts := make([][]int, len(words))
	for i := range counts {
		counts[i] = make([]int, 26)
	}

	for i, w := range words {
		for j := range w {
			scores[i] += score[getIndex(w[j])]
			counts[i][getIndex(w[j])] += 1
		}
	}

	totalCount := make([]int, 26)
	for _, l := range letters {
		totalCount[getIndex(l)] += 1
	}

	max := 0
	for i := 1; i < pow2(len(words)); i++ {
		n := i
		s := 0
		tempCount := make([]int, 26)

		for j := 0; j < len(words); j++ {
			if n%2 == 0 {
				n = n >> 1
				continue
			}

			enough := true
			for k := 0; k < 26; k++ {
				tempCount[k] += counts[j][k]
				if tempCount[k] > totalCount[k] {
					enough = false
					break
				}
			}
			if !enough {
				break
			}
			s += scores[j]
			n = n >> 1
		}

		if s > max {
			max = s
		}
	}

	return max
}

func getIndex(b byte) int {
	return int(b - byte('a'))
}

func pow2(n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r = r * 2
	}
	return r
}
