package _903

func numPermsDISequence(S string) int {
    command := []byte(S)
    mod := int(1e9 + 7)

    dp := make([][]int, len(S)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(S)+1)
        if i == 0 {
            for j := 0; j < len(dp[i]); j++ {
                dp[i][j] = 1
            }
        } else {
            left := len(dp) - i
            if command[i-1] == 'D' {
                //需要下降，那么j至少要有1才行
                for j := 1; j <= left; j++ {
                    for newJ := 0; newJ < j; newJ++ {
                        dp[i][newJ] += dp[i-1][j]
                        dp[i][newJ] %= mod
                    }
                }
            } else {
                //需要上升，那么j最多为剩余数量-1
                for j := 0; j < left; j++ {
                    for newJ := j; newJ < left; newJ++ {
                        dp[i][newJ] += dp[i-1][j]
                        dp[i][newJ] %= mod
                    }
                }
            }
        }
    }
    //fmt.Println(dp)
    return dp[len(S)][0]
}
