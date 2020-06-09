package _892

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func surfaceArea(grid [][]int) int {
	result := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				result += 2
				if i > 0 {
					result += max(grid[i][j]-grid[i-1][j], 0)
				} else {
					result += grid[i][j]
				}
				if j > 0 {
					result += max(grid[i][j]-grid[i][j-1], 0)
				} else {
					result += grid[i][j]
				}
				if i < len(grid)-1 {
					result += max(grid[i][j]-grid[i+1][j], 0)
				} else {
					result += grid[i][j]
				}
				if j < len(grid[i])-1 {
					result += max(grid[i][j]-grid[i][j+1], 0)
				} else {
					result += grid[i][j]
				}
			}
		}
	}
	return result
}
