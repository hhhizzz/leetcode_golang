package _562

func longestLine(M [][]int) int {
    result := 0
    if len(M) == 0 {
        return result
    }

    dp := make([][]int, len(M))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(M[i]))
    }
    //horizontal
    for i := 0; i < len(dp); i++ {
        for j := 0; j < len(dp[i]); j++ {
            if M[i][j] == 1 {
                if j == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = dp[i][j-1] + 1
                }
                if dp[i][j] > result {
                    result = dp[i][j]
                }
            }
        }
    }

    for i := 0; i < len(dp); i++ {
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = 0
        }
    }

    //vertical
    for j := 0; j < len(dp[0]); j++ {
        for i := 0; i < len(M); i++ {
            if M[i][j] == 1 {
                if i == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = dp[i-1][j] + 1
                }
                if dp[i][j] > result {
                    result = dp[i][j]
                }
            }
        }
    }

    for i := 0; i < len(dp); i++ {
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = 0
        }
    }
    //right anti-diagonal
    for i := 0; i < len(dp); i++ {
        for j := 0; j < len(dp[i]); j++ {
            if M[i][j] == 1 {
                if i == 0 || j == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = dp[i-1][j-1] + 1
                }
                if dp[i][j] > result {
                    result = dp[i][j]
                }
            }
        }
    }

    for i := 0; i < len(dp); i++ {
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = 0
        }
    }
    //left anti-diagonal
    for i := 0; i < len(dp); i++ {
        for j := len(dp[i]) - 1; j >= 0; j-- {
            if M[i][j] == 1 {
                if i == 0 || j == len(dp[i])-1 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = dp[i-1][j+1] + 1
                }
                if dp[i][j] > result {
                    result = dp[i][j]
                }
            }
        }
    }

    return result
}
