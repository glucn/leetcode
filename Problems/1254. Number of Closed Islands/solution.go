package solution

// Runtime 4 ms, faster than ?%
// Memory 4.9 MB, less than ?%

func closedIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if shouldVisit(i, j, grid, visited) && dfs(i, j, grid, visited) {
				count++
			}
		}
	}

	return count
}

var moves = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func dfs(i, j int, grid [][]int, visited [][]bool) bool {
	if visited[i][j] {
		return false
	}

	if i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1 {
		return false
	}

	visited[i][j] = true

	isClosed := true
	for _, m := range moves {
		if shouldVisit(i+m[0], j+m[1], grid, visited) {
			if !(dfs(i+m[0], j+m[1], grid, visited) || grid[i+m[0]][j+m[1]] == 1) {
				isClosed = false
			}
		}
	}
	return isClosed
}

func shouldVisit(i, j int, grid [][]int, visited [][]bool) bool {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 {
		return false
	}

	if visited[i][j] {
		return false
	}

	if grid[i][j] == 1 {
		return false
	}

	return true
}
