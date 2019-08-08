package _714

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxProfit(prices []int, fee int) int {
    hold := make([]int, len(prices))
    notHold := make([]int, len(prices))

    lastDay := len(prices) - 1

    hold[0] = -prices[0]
    notHold[0] = 0
    for day := 1; day <= lastDay; day++ {
        hold[day] = max(hold[day-1], notHold[day-1]-prices[day])
        notHold[day] = max(notHold[day-1], hold[day-1]+prices[day]-fee)
    }
    return notHold[lastDay]

}
