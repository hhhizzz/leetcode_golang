package _583

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//本体题可以转化为求word1和word2的最长匹配子序列
func minDistance(word1 string, word2 string) int {
	//dp[i][j]表示word1[i]和word2[j]之前最长能匹配的子序列
	dp := make([][]int, len(word1))
	length := 0
	for i := 0; i < len(word1); i++ {
		dp[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else {
				if i != 0 && j != 0 {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				} else if i != 0 {
					dp[i][j] = dp[i-1][j]
				} else if j != 0 {
					dp[i][j] = dp[i][j-1]
				}
			}
			length = max(length, dp[i][j])
		}
	}
	//fmt.Println(dp)
	//结果就是两个串长度相加减去最长匹配子序列的长度*2
	return len(word1) + len(word2) - length*2
}
