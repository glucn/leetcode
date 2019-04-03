package solution

// Runtime 108 ms
// Memory 7.8 MB

// TODO: there is a better solution

func shipWithinDays(weights []int, D int) int {
	sum := 0
	for _, w := range weights {
		sum += w
	}

	size := sum / D

	for {
		ship := make([]int, D)
		current := 0
		okay := true
		for i := 0; i < len(weights); i++ {
			if ship[current]+weights[i] <= size {
				ship[current] += weights[i]
			} else {
				if current == D-1 || weights[i] > size {
					okay = false
					break
				}
				current++
				ship[current] = weights[i]
			}
		}
		if !okay {
			size++
		} else {
			return size
		}
	}
}
