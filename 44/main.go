package _44

/*
动态规划处理，dp[i][j]表示模式串第i个是否能和源串第j个字符匹配，很容易想到当p[i]=='?'或者匹配的时候dp[i][j]==dp[i-1][j-1]

两个难点:
一个是处理边界条件，例如s或者p为空的时候，最简单的方法就是dp的时候算从1开始计数，这样不需要单独处理为空的情况了
第二个是p[i]=='*'的时候状态转移方程式 dp[i][j] = dp[i-1][j] || dp[i][j-1]，前一个表示*匹配空，后一个表示*匹配的长度能够扩展到当前s[j]字符

*/
func isMatch(s string, p string) bool {
	// 在前面加个空格更好理解
	s = " " + s
	p = " " + p
	dp := make([][]bool, len(p))
	for i := 0; i < len(p); i++ {
		dp[i] = make([]bool, len(s))
	}
	// 都为空显然可以匹配
	dp[0][0] = true
	// 看s为空时p是否能匹配
	for i := 1; i < len(p); i++ {
		if p[i] == '*' {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < len(p); i++ {
		for j := 1; j < len(s); j++ {
			if p[i] != '*' {
				if p[i] == s[j] || p[i] == '?' {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(p)-1][len(s)-1]
}
