package _198

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	dp := make([]int, len(nums))

	result := 0
	for i := 0; i < len(dp); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = max(dp[0], nums[i])
		} else {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
		result = max(result, dp[i])
	}
	return result
}
