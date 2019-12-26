package _115

func numDistinct(s string, t string) int {
    if len(t) == 0 {
        return 1
    }
    if len(s) < len(t) {
        return 0
    }

    arrayS := []byte(s)
    arrayT := []byte(t)

    dp := make([][]int, len(s))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(t))
    }
    for j := 0; j < len(t); j++ {
        count := 0
        for i := j; i < len(s); i++ {
            if arrayS[i] == arrayT[j] {
                if j == 0 {
                    count += 1
                } else {
                    count += dp[i-1][j-1]
                }
            }
            dp[i][j] = count
        }
    }
    //fmt.Println(dp)
    return dp[len(s)-1][len(t)-1]
}
