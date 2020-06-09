package _44

/*
动态规划处理，dp[i][j]表示模式串第i个是否能和源串第j个字符匹配，很容易想到当p[i]=='?'或者匹配的时候dp[i][j]==dp[i-1][j-1]

两个难点:
一个是处理边界条件，例如s或者p为空的时候，最简单的方法就是dp的时候算从1开始计数，这样不需要单独处理为空的情况了
第二个是p[i]=='*'的时候状态转移方程式 dp[i][j] = dp[i-1][j] || dp[i][j-1]，前一个表示*匹配空，后一个表示*匹配的长度能够扩展到当前s[j]字符

*/
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := 0; i <= len(p); i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	//都为空当然可以匹配
	dp[0][0] = true
	//处理s为空的情况 注意因为从1开始计数，注意取字符串的时候需要-1，像是p[i-1]和s[j-1]
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			if p[i-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else if p[i-1] == '?' || p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[len(p)][len(s)]
}
