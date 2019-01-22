package _63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	graph := make([][]int, m)
	obstacle := false
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
		if obstacle || obstacleGrid[i][0] == 1 {
			obstacle = true
			continue
		} else {
			graph[i][0] = 1
		}
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] != 1 {
			graph[0][j] = 1
		} else {
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				graph[i][j] = graph[i-1][j] + graph[i][j-1]
			}
		}
	}
	return graph[m-1][n-1]
}
