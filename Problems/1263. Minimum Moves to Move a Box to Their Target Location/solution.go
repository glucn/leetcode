package solution

import "math"

// Runtime 1440 ms, faster than ?%
// Memory 16.7 MB, less than ?%

func minPushBox(grid [][]byte) int {
	visited := make(map[game]int)

	initial := game{}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == byte('B') {
				initial.xBox, initial.yBox = i, j
			}
			if grid[i][j] == byte('S') {
				initial.xPlayer, initial.yPlayer = i, j
			}
		}
	}

	queue := []solution{{game: initial, step: 0}}
	visited[initial] = 0
	minStep := math.MaxInt32

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		for _, m := range moves {
			gg, ok, isBoxMoved, endgame := nextGame(s.game, m[0], m[1], grid)
			if !ok {
				continue
			}

			if endgame {
				minStep = min(minStep, s.step+1)
			}

			nextStep := s.step
			if isBoxMoved {
				nextStep = s.step + 1
			}

			if _, ok := visited[gg]; ok && visited[gg] <= nextStep {
				continue
			}

			queue = append(queue, solution{game: gg, step: nextStep})
			visited[gg] = nextStep
		}
	}
	if minStep == math.MaxInt32 {
		return -1
	}
	return minStep
}

var moves = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func nextGame(g game, x, y int, grid [][]byte) (gg game, ok bool, isBoxMoved bool, endgame bool) {
	nextPX := g.xPlayer + x
	nextPY := g.yPlayer + y

	if nextPX < 0 || nextPX >= len(grid) || nextPY < 0 || nextPY >= len(grid[0]) {
		return game{}, false, false, false
	}

	if grid[nextPX][nextPY] == byte('#') {
		return game{}, false, false, false
	}

	if nextPX == g.xBox && nextPY == g.yBox {
		nextBX := g.xBox + x
		nextBY := g.yBox + y
		if nextBX < 0 || nextBX >= len(grid) || nextBY < 0 || nextBY >= len(grid[0]) {
			return game{}, false, false, false
		}

		if grid[nextBX][nextBY] == byte('#') {
			return game{}, false, false, false
		}

		if grid[nextBX][nextBY] == byte('T') {
			// endgame
			return game{}, true, true, true
		}
		return game{
			xBox:    nextBX,
			yBox:    nextBY,
			xPlayer: nextPX,
			yPlayer: nextPY,
		}, true, true, false
	}

	return game{
		xBox:    g.xBox,
		yBox:    g.yBox,
		xPlayer: nextPX,
		yPlayer: nextPY,
	}, true, false, false
}

type game struct {
	xBox, yBox       int
	xPlayer, yPlayer int
}

type solution struct {
	game game
	step int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
