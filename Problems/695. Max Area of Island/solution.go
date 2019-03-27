package solution

// Runtime 16 ms
// Memory 4.7 MB

func maxAreaOfIsland(grid [][]int) int {
    touched := make([][]bool, len(grid))
    for i := range grid {
        touched[i] = make([]bool, len(grid[i]))
    }
    
    maxArea := 0

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 0 || touched[i][j] {
                continue
            }
            area := traverse(i,j, grid, touched)
            if area > maxArea {
                maxArea = area
            }
        }
    }
    
    return maxArea
    
}

func traverse(x, y int, grid [][]int, touched [][]bool) int {
    
    if grid[x][y] == 0 || touched[x][y] {
        return 0
    }
    
    touched[x][y] = true
    area := 1
    
    if x > 0  {
        area += traverse(x-1, y, grid, touched)
    }
    if x < len(grid) - 1 {
        area += traverse(x+1, y, grid, touched)
    }
    if y > 0 {
        area += traverse(x, y-1, grid, touched)
    }
    if y < len(grid[x]) - 1 {
        area += traverse(x, y+1, grid, touched)
    }
    return area
    
}