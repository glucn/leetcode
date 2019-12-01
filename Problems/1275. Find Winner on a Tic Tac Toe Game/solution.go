package solution

// Runtime 0 ms, faster than ?%
// Memory 2 MB, less than ?%

func tictactoe(moves [][]int) string {
	rows := make([]int, 3)
	cols := make([]int, 3)
	leftToRightDiagonal := 0 // topleft to bottomright
	rightToLeftDiagonal := 0 // topright to bottomleft

	for i, move := range moves {
		m := 1 // A: 1; B: -1
		if i%2 == 1 {
			m = -1
		}

		rows[move[0]] += m
		if rows[move[0]] == 3*m {
			return win(i)
		}

		cols[move[1]] += m
		if cols[move[1]] == 3*m {
			return win(i)
		}

		if move[0] == move[1] {
			leftToRightDiagonal += m
			if leftToRightDiagonal == 3*m {
				return win(i)
			}
		}

		if move[0]+move[1] == 2 {
			rightToLeftDiagonal += m
			if rightToLeftDiagonal == 3*m {
				return win(i)
			}
		}

	}

	if len(moves) < 9 {
		return "Pending"
	}

	return "Draw"
}

func win(i int) string {
	if i%2 == 0 {
		return "A"
	}
	return "B"
}
