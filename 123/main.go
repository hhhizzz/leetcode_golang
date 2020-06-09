package _123

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	//dp[i][j]表示第i天执行了j次操作的最大利润，prices[0]表示第1天的价格
	var dp [][]int
	//买入或卖出算一次操作，因此总共可以执行4次
	action := 4

	lastDay := len(prices)

	dp = make([][]int, lastDay+1)
	for i := 0; i <= lastDay; i++ {
		dp[i] = make([]int, action+1)
	}
	//第0天0次操作获利是0，其余都是负无穷，这样能保证某天买操作能够执行
	//例如第1次买操作，如果第0天1次操作的获利是0，那么dp后认为不应该买，因为买了获利为负，而不买0的获利更高一点
	//那么把第0天n次操作的获利调为负无穷即可
	for times := 1; times < action; times++ {
		dp[0][times] = math.MinInt32
	}

	/*
	   状态转移方程：
	    当j为奇数 表示当前持有股票，可以卖出或者不卖
	    dp[i+1][j] = max(dp[i][j-1]+prices[i],dp[i][j])
	    当j为偶数 表示当前没有持有股票，可以买或者不买
	    dp[i+1][j] = max(dp[i][j-1]-prices[i],dp[i][j])
	*/

	result := 0
	for times := 1; times <= action; times++ {
		for day := 1; day <= lastDay; day++ {
			if times%2 == 1 {
				dp[day][times] = max(dp[day-1][times-1]-prices[day-1], dp[day-1][times])
			} else {
				dp[day][times] = max(dp[day-1][times-1]+prices[day-1], dp[day-1][times])
			}
			result = max(dp[day][times], result)
		}
	}
	fmt.Println(dp)

	return result
}
