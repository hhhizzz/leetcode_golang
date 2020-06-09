package _673

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	number := make([]int, len(nums))
	for i := 0; i < len(number); i++ {
		number[i] = 1
	}
	longest := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					number[i] = number[j]
				} else if dp[j]+1 == dp[i] {
					number[i] += number[j]
				}
			}
		}
		longest = max(longest, dp[i])
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == longest {
			result += number[i]
		}
	}
	return result
}
