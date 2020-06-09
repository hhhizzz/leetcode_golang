package _887

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for k := range dp {
		dp[k] = make([]int, N+1)
		if k != 0 {
			dp[k][1] = 1
		}
	}
	for m := 2; m <= N; m++ {
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k-1][m-1] + 1 + dp[k][m-1]
			if dp[k][m] >= N {
				return m
			}
		}
	}
	return -1
}

//方法一：正向动态规划
func superEggDrop1(K int, N int) int {
	dp := make([][]int, K+1)
	for k := range dp {
		dp[k] = make([]int, N+1)
	}
	for k := 1; k <= K; k++ {
		dp[k][1] = 1
	}
	for n := 1; n <= N; n++ {
		dp[1][n] = n
	}
	for k := 2; k <= K; k++ {
		x := 1
		for n := 2; n <= N; n++ {
			for ; x < n && max(dp[k-1][x-1], dp[k][n-x]) > max(dp[k-1][x], dp[k][n-x-1]); x++ {

			}
			dp[k][n] = max(dp[k-1][x-1], dp[k][n-x]) + 1
		}
	}
	return dp[K][N]
}
