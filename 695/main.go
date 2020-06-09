package _695

type pos struct {
	x int
	y int
}

func dfs(grid [][]int, visited map[pos]bool, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	if _, ok := visited[pos{i, j}]; ok {
		return 0
	}
	visited[pos{i, j}] = true
	left := dfs(grid, visited, i, j-1)
	up := dfs(grid, visited, i-1, j)
	right := dfs(grid, visited, i, j+1)
	down := dfs(grid, visited, i+1, j)
	return left + up + right + down + 1
}

func maxAreaOfIsland(grid [][]int) int {
	visited := map[pos]bool{}
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if _, ok := visited[pos{i, j}]; !ok && grid[i][j] == 1 {
				current := dfs(grid, visited, i, j)
				if current > result {
					result = current
				}
			}
		}
	}
	return result
}
