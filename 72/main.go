package _72

func min2(i, j int) int {
    if i < j {
        return i
    }
    return j
}

func min(i, j, k int) int {
    return min2(min2(i, j), k)
}

func minDistance(word1 string, word2 string) int {
    if len(word1) == 0 {
        return len(word2)
    }
    if len(word2) == 0 {
        return len(word1)
    }
    dp := make([][]int, len(word1)+1)
    dp[0] = make([]int, len(word2)+1)
    for i := 0; i <= len(word1); i++ {
        dp[i] = make([]int, len(word2)+1)
        dp[i][0] = i
    }
    for j := 0; j <= len(word2); j++ {
        dp[0][j] = j
    }
    for i := 1; i <= len(word1); i++ {
        ci := word1[i-1]
        for j := 1; j <= len(word2); j++ {
            cj := word2[j-1]
            if ci == cj {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }
    return dp[len(word1)][len(word2)]
}
