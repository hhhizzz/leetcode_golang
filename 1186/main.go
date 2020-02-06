package _1186

import "math"

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
    result := math.MinInt32
    for i := 0; i < len(arr); i++ {
        if i == 0 {
            dp[i] = arr[i]
            if arr[i] > 0 {
                dp2[i] = arr[i]
            }
            result = max(result, dp[i])
        } else {
            dp[i] = max(arr[i], dp[i-1]+arr[i])
            dp2[i] = max(dp2[i-1]+arr[i], dp[i-1])
            result = max3(result, dp[i], dp2[i])
        }
    }
    return result
}
