package _548

func splitArray(nums []int) bool {
    if len(nums) < 7 {
        return false
    }
    dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++ {
        if i == 0 {
            dp[i] = nums[i]
        } else {
            dp[i] = nums[i] + dp[i-1]
        }
    }
    for i := 1; i < len(nums)-5; i++ {
        for k := i + 4; k < len(nums)-1; k++ {
            if dp[i-1] == dp[len(dp)-1]-dp[k] {
                for j := i + 2; j < k-1; j++ {
                    if dp[j-1]-dp[i] == dp[i-1] && dp[k-1]-dp[j] == dp[i-1] {
                        return true
                    }
                }
            }
        }
    }
    return false
}
