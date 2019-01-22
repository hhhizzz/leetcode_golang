package _62

func uniquePaths(m int, n int) int {
	var graph = make([][]int, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
		graph[i][0] = 1
	}
	for j := 0; j < n; j++ {
		graph[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			graph[i][j] = graph[i-1][j] + graph[i][j-1]
		}
	}
	return graph[m-1][n-1]
}
