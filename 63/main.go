package _63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if obstacleGrid[i][j] == 0 {
				var current int
				if i != 0 {
					current += dp[i-1][j]
				}
				if j != 0 {
					current += dp[i][j-1]
				}
				if i != 0 || j != 0 {
					dp[i][j] = current
				}
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
