package _256

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minCost(costs [][]int) int {
    if len(costs) == 0 {
        return 0
    }
    dp := [][]int{costs[0], {0, 0, 0}}
    for i := 1; i < len(costs); i++ {
        last := dp[(i-1)%2]
        current := dp[i%2]
        current[0] = min(last[1], last[2]) + costs[i][0]
        current[1] = min(last[0], last[2]) + costs[i][1]
        current[2] = min(last[0], last[1]) + costs[i][2]
    }
    last := dp[(len(costs)-1)%2]
    return min(min(last[0], last[1]), last[2])
}
