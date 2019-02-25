package solution

// Runtime 0 ms
// Memory 1.9 MB

func numRookCaptures(board [][]byte) int {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'R' {
				return up(board, i, j) + down(board, i, j) + left(board, i, j) + right(board, i, j)
			}
		}
	}
	return -1
}

func up(board [][]byte, i, j int) int {
	if i == 0 {
		return 0
	}
	if board[i-1][j] == 'p' {
		return 1
	}
	if board[i-1][j] == '.' {
		return up(board, i-1, j)
	}
	return 0
}

func down(board [][]byte, i, j int) int {
	if i == len(board)-1 {
		return 0
	}
	if board[i+1][j] == 'p' {
		return 1
	}
	if board[i+1][j] == '.' {
		return down(board, i+1, j)
	}
	return 0
}

func left(board [][]byte, i, j int) int {
	if j == 0 {
		return 0
	}
	if board[i][j-1] == 'p' {
		return 1
	}
	if board[i][j-1] == '.' {
		return left(board, i, j-1)
	}
	return 0
}

func right(board [][]byte, i, j int) int {
	if j == len(board[i])-1 {
		return 0
	}
	if board[i][j+1] == 'p' {
		return 1
	}
	if board[i][j+1] == '.' {
		return right(board, i, j+1)
	}
	return 0
}
