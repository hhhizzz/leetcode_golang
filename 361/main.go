package _361

func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	bomb := make([][]int, m)
	for i := 0; i < m; i++ {
		bomb[i] = make([]int, n)
	}
	result := 0

	for i := 0; i < m; i++ {
		//to right
		current := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == '0' {
				bomb[i][j] += current
				if bomb[i][j] > result {
					result = bomb[i][j]
				}
			} else if grid[i][j] == 'E' {
				current++
			} else if grid[i][j] == 'W' {
				current = 0
			}
		}
		//to left
		current = 0
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				bomb[i][j] += current
				if bomb[i][j] > result {
					result = bomb[i][j]
				}
			} else if grid[i][j] == 'E' {
				current++
			} else if grid[i][j] == 'W' {
				current = 0
			}
		}
	}
	for j := 0; j < n; j++ {
		//to up
		current := 0
		for i := 0; i < m; i++ {
			if grid[i][j] == '0' {
				bomb[i][j] += current
				if bomb[i][j] > result {
					result = bomb[i][j]
				}
			} else if grid[i][j] == 'E' {
				current++
			} else if grid[i][j] == 'W' {
				current = 0
			}
		}
		//to down
		current = 0
		for i := m - 1; i >= 0; i-- {
			if grid[i][j] == '0' {
				bomb[i][j] += current
				if bomb[i][j] > result {
					result = bomb[i][j]
				}
			} else if grid[i][j] == 'E' {
				current++
			} else if grid[i][j] == 'W' {
				current = 0
			}
		}
	}

	return result
}
