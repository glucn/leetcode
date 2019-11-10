package _733__Flood_Fill

// Runtime 12 ms, faster than 35.64%
// Memory 6.3 MB, less than 100%

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visited := make([][]bool, len(image))
	for i := range visited {
		visited[i] = make([]bool, len(image[0]))
	}

	dfs(sr, sc, image, visited, image[sr][sc], newColor)
	return image
}

var moves = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func dfs(i, j int, image [][]int, visited [][]bool, oldColor, newColor int) {
	if i < 0 || i > len(image)-1 || j < 0 || j > len(image[0])-1 {
		return
	}

	if visited[i][j] {
		return
	}

	if image[i][j] != oldColor {
		return
	}

	image[i][j] = newColor
	visited[i][j] = true

	for _, m := range moves {
		dfs(i+m[0], j+m[1], image, visited, oldColor, newColor)
	}
}
