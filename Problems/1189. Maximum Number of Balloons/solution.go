package solution

// Runtime 0 ms, faster than 100%
// Memory 2.2 MB, less than 100%

var balloon = []rune{'b', 'a', 'l', 'o', 'n'}

func maxNumberOfBalloons(text string) int {
	count := make(map[rune]int)

	for _, b := range balloon {
		count[b] = 0
	}

	for _, c := range text {
		if inSlice(c, balloon) {
			count[c] = count[c] + 1
		}
	}

	count['l'] = count['l'] >> 1
	count['o'] = count['o'] >> 1

	min := len(text)
	for _, v := range count {
		if v < min {
			min = v
		}
	}
	return min
}

func inSlice(r rune, runes []rune) bool {
	for _, rr := range runes {
		if r == rr {
			return true
		}
	}
	return false
}
