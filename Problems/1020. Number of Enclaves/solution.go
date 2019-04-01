package solution

// Runtime 60 ms
// Memory 6.6 MB

func numEnclaves(A [][]int) int {
	check := make([][]bool, len(A))
	total := 0

	for i := 0; i < len(A); i++ {
		check[i] = make([]bool, len(A[0]))
		for j := 0; j < len(A[0]); j++ {
			check[i][j] = true
		}
	}

	for i := 0; i < len(A); i++ {
		if A[i][0] == 1 && check[i][0] {
			markFalse(A, check, i, 0)
		}
		if A[i][len(A[0])-1] == 1 && check[i][len(A[0])-1] {
			markFalse(A, check, i, len(A[0])-1)
		}
	}
	for j := 0; j < len(A[0]); j++ {
		if A[0][j] == 1 && check[0][j] {
			markFalse(A, check, 0, j)
		}
		if A[len(A)-1][j] == 1 && check[len(A)-1][j] {
			markFalse(A, check, len(A)-1, j)
		}
	}

	for i := len(A) - 1; i >= 0; i-- {
		for j := len(A[i]) - 1; j >= 0; j-- {
			if A[i][j] == 1 && check[i][j] {
				total++
			}
		}
	}

	return total
}

func markFalse(A [][]int, check [][]bool, i, j int) {
	check[i][j] = false
	if i != 0 && A[i-1][j] == 1 && check[i-1][j] {
		markFalse(A, check, i-1, j)
	}
	if i != len(A)-1 && A[i+1][j] == 1 && check[i+1][j] {
		markFalse(A, check, i+1, j)
	}
	if j != 0 && A[i][j-1] == 1 && check[i][j-1] {
		markFalse(A, check, i, j-1)
	}
	if j != len(A[i])-1 && A[i][j+1] == 1 && check[i][j+1] {
		markFalse(A, check, i, j+1)
	}
}
