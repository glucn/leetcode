package solution

// Runtime 0 ms
// Memory 4.1 MB

func equalSubstring(s string, t string, maxCost int) int {
	diff := make([]int, len(s))
	for i := range s {
		diff[i] = abs(int(s[i]) - int(t[i]))
	}

	cost := 0
	i := 0
	j := 0
	for ; j < len(s); j++ {
		cost += diff[j]
		if cost > maxCost {
			cost -= diff[i]
			i++
		}
	}

	return j - i
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
