package solution

// Runtime 0 ms
// Memory 2.2 MB

func transformArray(arr []int) []int {
	if len(arr) <= 2 {
		return arr
	}
	temp := make([][]int, 2)
	temp[0] = arr
	temp[1] = make([]int, len(arr))
	temp[1][0] = arr[0]
	temp[1][len(arr)-1] = arr[len(arr)-1]

	i := 1
	for !shouldStop(temp[0], temp[1]) {
		for j := 1; j < len(arr)-1; j++ {
			if temp[(i-1)%2][j] < temp[(i-1)%2][j-1] && temp[(i-1)%2][j] < temp[(i-1)%2][j+1] {
				temp[i%2][j] = temp[(i-1)%2][j] + 1
			} else if temp[(i-1)%2][j] > temp[(i-1)%2][j-1] && temp[(i-1)%2][j] > temp[(i-1)%2][j+1] {
				temp[i%2][j] = temp[(i-1)%2][j] - 1
			} else {
				temp[i%2][j] = temp[(i-1)%2][j]
			}
		}
		// fmt.Println(temp[i%2])
		i++
	}
	return temp[i%2]
}

func shouldStop(a1, a2 []int) bool {
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
