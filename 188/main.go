package _188

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxProfit(k int, prices []int) int {
    if len(prices) == 0 || k == 0 {
        return 0
    }
    lastDay := len(prices) - 1
    mp := make([][]int, len(prices))

    result := 0
    k *= 2
    if k >= len(prices) {
        for day := 0; day < len(prices)-1; day++ {
            if prices[day+1] > prices[day] {
                result += prices[day+1] - prices[day]
            }
        }
        return result
    }

    for day := 0; day <= lastDay; day++ {
        mp[day] = make([]int, k+1)
    }

    mp[0][0] = 0
    mp[0][1] = -prices[0]

    for times := 2; times <= k; times++ {
        mp[0][times] = -int(^uint(0) >> 1) - 1
    }

    for times := 1; times <= k; times++ {
        for day := 1; day <= lastDay; day++ {
            if times%2 == 1 {
                mp[day][times] = max(mp[day-1][times], mp[day][times-1]-prices[day])
            } else {
                mp[day][times] = max(mp[day-1][times], mp[day][times-1]+prices[day])
            }
            if mp[day][times] > result {
                result = mp[day][times]
            }
        }
    }

    return result
}
