package solution

// Runtime 2876 ms  // TODO: find a better solution
// Memory 9. MB

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, q := range queries {
		if isLit(lamps, q[0], q[1]) {
			result[i] = 1
		}
		turnOff(lamps, q[0], q[1])
	}

	return result

}

func isLit(lamps [][]int, x, y int) bool {
	for _, l := range lamps {
		if l[0] == x {
			return true
		}
		if l[1] == y {
			return true
		}
		if (x - l[0]) == (y - l[1]) {
			return true
		}
		if (x - l[0]) == (l[1] - y) {
			return true
		}
	}
	return false
}

func turnOff(lamps [][]int, x, y int) {
	var toBeDeleted []int
	for i, l := range lamps {
		if l[0] >= x-1 && l[0] <= x+1 && l[1] >= y-1 && l[1] <= y+1 {
			toBeDeleted = append(toBeDeleted, i)
		}
	}

	for _, t := range toBeDeleted {
		lamps[t] = lamps[len(lamps)-1]
	}
	lamps = lamps[:len(lamps)-len(toBeDeleted)]

}
