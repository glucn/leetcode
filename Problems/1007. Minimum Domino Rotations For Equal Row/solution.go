package solution

// Runtime 368 ms
// Memory 9.7 MB

func minDominoRotations(A []int, B []int) int {
	setUp := make([]map[int]struct{}, 7)
	setDown := make([]map[int]struct{}, 7)

	for i := 0; i < 7; i++ {
		setUp[i] = make(map[int]struct{})
		setDown[i] = make(map[int]struct{})
	}

	for i, a := range A {
		setUp[a][i] = struct{}{}
	}
	for i, b := range B {
		setDown[b][i] = struct{}{}
	}

	min := len(A)
	for i := 1; i <= 6; i++ {
		if len(setUp[i])+len(setDown[i]) < len(A) {
			continue
		}

		merge := mergeMap(setUp[i], setDown[i])
		if len(merge) == len(A) {
			onlyInUp := findDistinct(setUp[i], setDown[i])
			onlyInDown := findDistinct(setDown[i], setUp[i])
			if onlyInUp <= onlyInDown {
				if onlyInUp < min {
					min = onlyInUp
				}
			} else {
				if onlyInDown < min {
					min = onlyInDown
				}
			}
		}
	}
	if min == len(A) {
		return -1
	}
	return min
}

func mergeMap(a, b map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}
	return result
}

func findDistinct(a, b map[int]struct{}) int {
	count := 0
	for k := range b {
		if _, ok := a[k]; !ok {
			count++
		}
	}
	return count
}
