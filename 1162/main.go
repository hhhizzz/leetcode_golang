package _1162

func maxDistance(grid [][]int) int {
	var q [][2]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
			}
		}
	}
	if len(q) == 0 || len(q) == len(grid)*len(grid[0]) {
		return -1
	}
	result := 0
	for len(q) != 0 {
		current := q[0]
		q = q[1:]
		i := current[0]
		j := current[1]
		if i > 0 && grid[i-1][j] == 0 {
			grid[i-1][j] = grid[i][j] + 1
			q = append(q, [2]int{i - 1, j})
		}
		if j > 0 && grid[i][j-1] == 0 {
			grid[i][j-1] = grid[i][j] + 1
			q = append(q, [2]int{i, j - 1})
		}
		if i < len(grid)-1 && grid[i+1][j] == 0 {
			grid[i+1][j] = grid[i][j] + 1
			q = append(q, [2]int{i + 1, j})
		}
		if j < len(grid[0])-1 && grid[i][j+1] == 0 {
			grid[i][j+1] = grid[i][j] + 1
			q = append(q, [2]int{i, j + 1})
		}
		if grid[i][j] > result {
			result = grid[i][j]
		}
	}
	return result - 1
}
