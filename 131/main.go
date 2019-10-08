package _131

var dp [][]bool

func solve(start int, s string) [][]string {
    var result [][]string

    if start == len(s) {
        result = append(result, []string{})
        return result
    }

    for i := start; i < len(s); i++ {
        if !dp[start][i] {
            continue
        }
        next := solve(i+1, s)
        for n := range next {
            next[n] = append([]string{s[start : i+1]}, next[n]...)
            result = append(result, next[n])
        }
    }
    return result
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
    return solve(0, s)
}
