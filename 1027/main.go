package _1027

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestArithSeqLength(A []int) int {
	dp := make([]map[int]int, len(A))
	result := 0
	for i := 0; i < len(A); i++ {
		dp[i] = map[int]int{}
		for j := i - 1; j >= 0; j-- {
			diff := A[i] - A[j]
			dp[i][diff] = max(dp[j][diff]+1, dp[i][diff])
			result = max(dp[i][diff]+1, result)
		}
	}
	return result
}
