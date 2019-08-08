package _309

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    mp := make([][2]int, len(prices))
    lastDay := len(prices) - 1

    //0: sale 1: buy

    mp[0][0] = 0
    mp[0][1] = -prices[0]

    mp[1][0] = max(0, prices[1]-prices[0])
    mp[1][1] = max(-prices[0], -prices[1])

    for day := 2; day <= lastDay; day++ {
        mp[day][0] = max(mp[day-1][0], mp[day-1][1]+prices[day])
        mp[day][1] = max(mp[day-2][0]-prices[day], mp[day-1][1])
    }
    return max(mp[lastDay][0], mp[lastDay][1])
}
