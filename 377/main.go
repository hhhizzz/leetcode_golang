package _377

func combinationSum4(nums []int, target int) int {
	//dp[i]表示加到i有多少种可能，dp[i] = sum(dp[i-num]) 其中num是nums中的数
	dp := make([]int, target+1)
	for current := 1; current <= target; current++ {
		for _, num := range nums {
			if num == current {
				dp[current] += 1
			} else if num < current {
				dp[current] += dp[current-num]
			}
		}
	}
	return dp[target]
}
