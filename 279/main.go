package _279

import "math"

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := 1; i < len(dp); i++ {
        dp[i] = math.MaxInt32
        for j := 1; j*j <= i; j++ {
            dp[i] = min(dp[i], dp[i-j*j]+1)
        }
    }
    return dp[n]
}
