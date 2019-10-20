package solution

// Runtime 4 ms
// Memory 3.7 MB

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	for i := 2; i < len(coordinates); i++ {
		if (coordinates[1][1]-coordinates[0][1])*(coordinates[i][0]-coordinates[0][0]) != (coordinates[i][1]-coordinates[0][1])*(coordinates[1][0]-coordinates[0][0]) {
			return false
		}
	}
	return true
}
