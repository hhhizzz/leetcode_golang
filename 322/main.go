package _322

import "math"

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = math.MaxInt32
        for j := 0; j < len(coins); j++ {
            if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
                dp[i] = min(dp[i], dp[i-coins[j]]+1)
            }
        }
        if dp[i] == math.MaxInt32 {
            dp[i] = -1
        }
    }
    return dp[amount]
}
