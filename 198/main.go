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
    for i := range nums {
        dp[i] = nums[i]
        for j := i - 2; j >= 0; j-- {
            dp[i] = max(dp[i], dp[j]+nums[i])
        }
        result = max(result, dp[i])
    }
    return result
}
