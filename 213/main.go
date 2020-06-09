package _213

//可以用线段树或者树状数组优化，但这道题数据太小，没必要
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		max := 0
		for j := i - 2; j >= 0; j-- {
			if dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + nums[i]
		if dp[i] > result {
			result = dp[i]
		}
	}
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		max := 0
		for j := i - 2; j >= 0; j-- {
			if dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + nums[i]
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
