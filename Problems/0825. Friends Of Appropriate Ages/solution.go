package solution

// Runtime 32 ms
// Memory 6.2 MB

func numFriendRequests(ages []int) int {
	ageCount := make([]int, 121)
	for _, a := range ages {
		ageCount[a]++
	}

	count := 0
	for i := 0; i < len(ageCount); i++ {
		for j := 0; j < len(ageCount); j++ {

			if !(j > 100 && i < 100) && j > i/2+7 && j <= i {
				if i != j {
					count += ageCount[i] * ageCount[j]
				} else {
					count += ageCount[i] * (ageCount[i] - 1)
				}

			}
		}
	}
	return count
}
