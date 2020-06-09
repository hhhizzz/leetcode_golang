package _264

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	two := 0
	three := 0
	five := 0
	for i := 1; i < n; i++ {
		dp[i] = min3(dp[two]*2, dp[three]*3, dp[five]*5)
		if dp[i] >= dp[two]*2 {
			two++
		}
		if dp[i] >= dp[three]*3 {
			three++
		}
		if dp[i] >= dp[five]*5 {
			five++
		}
		//fmt.Println(dp[:i+1])
	}
	return dp[n-1]
}
