package _131

var dp [][]bool

func dfs(start int, temp *[]string, result *[][]string, s string) {
    if start == len(s) {
        newTemp := make([]string, len(*temp))
        copy(newTemp, *temp)
        *result = append(*result, newTemp)
        return
    }
    for i := start; i < len(s); i++ {
        if dp[start][i] {
            *temp = append(*temp, s[start:i+1])
            dfs(i+1, temp, result, s)
            *temp = (*temp)[:len(*temp)-1]
        }
    }
}

func partition(s string) [][]string {
    //dp[i][j]表示s[i]~s[j]是回文串
    dp = make([][]bool, len(s))
    for i := range dp {
        dp[i] = make([]bool, len(s))
    }

    for length := 1; length <= len(s); length++ {
        for i := 0; i < len(s); i++ {
            j := i + length - 1
            if j < len(s) && s[i] == s[j] {
                if length <= 3 || dp[i+1][j-1] {
                    dp[i][j] = true
                }
            }
        }
    }
    var result [][]string
    var temp []string

    dfs(0, &temp, &result, s)
    return result
}
