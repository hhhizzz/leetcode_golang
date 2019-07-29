package _122

func maxProfit(prices []int) int {
    result := 0
    for i := 0; i < len(prices)-1; i++ {
        if prices[i+1] > prices[i] {
            result += prices[i+1] - prices[i]
        }
    }
    return result
}
