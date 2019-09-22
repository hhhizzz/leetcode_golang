package _887

func superEggDrop(K int, N int) int {
    dp := make([][]int, K+1)
    for k := range dp {
        dp[k] = make([]int, N+1)
        if k > 0 {
            dp[k][1] = 1
        }
    }
    for m := 1; m <= N; m++ {
        for k := 1; k <= K; k++ {
            dp[k][m] = dp[k-1][m-1] + 1 + dp[k][m-1]
            if dp[k][m] >= N {
                return m
            }
        }
    }
    return N
}
