package _983

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func min3(a, b, c int) int {
    return min(min(a, b), c)
}

func mincostTickets(days []int, costs []int) int {
    dp := make([]int, len(days)+1)
    for i := len(days) - 1; i >= 0; i-- {
        day := days[i]
        dp[i] = costs[0] + dp[i+1]
        j := i
        for j < len(days) && days[j] < day+7 {
            j++
        }
        dp[i] = min(dp[i], dp[j]+costs[1])
        for j < len(days) && days[j] < day+30 {
            j++
        }
        dp[i] = min(dp[i], dp[j]+costs[2])
    }
    return dp[0]
}
