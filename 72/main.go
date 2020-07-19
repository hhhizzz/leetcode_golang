package _72

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 动态规划 dp[i][j] 表示word1的前i个字符转换到word2的前j个字符最少需要的次数
func minDistance(word1 string, word2 string) int {
	// 字符串动规题在前面垫一个空格会更方便操作
	word1 = " " + word1
	word2 = " " + word2
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		// 初始化（下同） 很显然 零个字符要变成i个或者i个字符变成0个最快也需要i次
		dp[i][0] = i
	}
	for j := 0; j < m; j++ {
		dp[0][j] = j
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			/*
				考虑得到dp[i][j]有3种情况
				1. word1增加一个字符或者word2减少一个字符，此时利用word1[:i]变换到word2[:j+1]的次数加一即可也就是dp[i-1][j]+1
				2. word2减少一个字符或者word1增加一个字符，对应dp[i][j-1]+1
				3. word1变换一个字符得到word[2] 此时看word1[i]与word2[j]是否匹配，如果匹配，那么也不需要变换，直接就是dp[i-1][j-1]，否则还需要+1
			*/
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			if word1[i] == word2[j] {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j])
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i][j])
			}
		}
	}
	return dp[n-1][m-1]
}
