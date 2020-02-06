package _1186

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func max3(a, b, c int) int {
    return max(max(a, b), c)
}

func maximumSum(arr []int) int {
    dp := make([]int, len(arr))
    dp2 := make([]int, len(arr))
    dp[0] = arr[0]
    dp2[0] = arr[0]
    result := arr[0]
    for i := 1; i < len(arr); i++ {
        dp[i] = max(dp[i-1]+arr[i], arr[i])
        dp2[i] = max(dp2[i-1]+arr[i], dp[i-1])
        result = max3(result, dp[i], dp2[i])
    }
    return result
}
