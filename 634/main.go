package _634

func findDerangement(n int) int {
    dp := [3]int{0, 1, 2}
    mod := int(1e9 + 7)
    for i := 3; i < n; i++ {
        dp[i%3] = (dp[(i-1)%3] + dp[(i-2)%3]) * i % mod
    }
    return dp[(n-1)%3]
}
