package _62

func uniquePaths(m int, n int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := 0; i < m; i++ {
		current := i % 2
		last := (current + 1) % 2
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[current][j] = dp[last][j]
			} else {
				dp[current][j] = dp[last][j] + dp[current][j-1]
			}
		}
	}
	current := (m - 1) % 2
	return dp[current][n-1]
}
