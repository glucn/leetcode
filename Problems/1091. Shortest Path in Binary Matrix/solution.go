package solution

// Runtime 48 ms
// Memory 6.5 MB

type cell struct {
	x, y   int
	length int
}

var moves = []cell{
	{x: 1, y: 0},
	{x: -1, y: 0},
	{x: 0, y: 1},
	{x: 0, y: -1},
	{x: 1, y: 1},
	{x: 1, y: -1},
	{x: -1, y: 1},
	{x: -1, y: -1},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	var queue []cell
	queue = append(queue, cell{x: 0, y: 0, length: 1})

	for len(queue) > 0 {
		g := queue[0]
		queue = queue[1:]

		if g.x == len(grid)-1 && g.y == len(grid[0])-1 {
			return g.length
		}

		for _, m := range moves {
			if shouldVisit(g.x+m.x, g.y+m.y, grid, visited) {
				queue = append(queue, cell{x: g.x + m.x, y: g.y + m.y, length: g.length + 1})
				visited[g.x+m.x][g.y+m.y] = true
			}
		}
	}

	return -1
}

func shouldVisit(x, y int, grid [][]int, visited [][]bool) bool {
	if x < 0 || x > len(grid)-1 {
		return false
	}
	if y < 0 || y > len(grid[0])-1 {
		return false
	}
	if grid[x][y] == 1 {
		return false
	}
	if visited[x][y] {
		return false
	}
	return true
}
