package _121

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buyPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		buyPrice = min(buyPrice, prices[i])
		profit = max(profit, prices[i]-buyPrice)
	}
	return profit
}
