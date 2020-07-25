package _97

//方法一：直接深搜
func isInterleave1(s1 string, s2 string, s3 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	n3 := len(s3)
	if n1+n2 != n3 {
		return false
	}
	if n1 == 0 {
		return s2 == s3
	}
	if n2 == 0 {
		return s1 == s3
	}

	var res bool
	if s3[0] == s1[0] {
		res = isInterleave(s1[1:], s2, s3[1:])
		if res {
			return true
		}
	}
	if s3[0] == s2[0] {
		res = isInterleave(s1, s2[1:], s3[1:])
	}
	return res
}

//方法二：动态规划
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	n3 := len(s3)
	if n1+n2 != n3 {
		return false
	}

	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true

	for i := 1; i <= n1; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = true
		} else {
			break
		}
	}

	for j := 1; j <= n2; j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = true
		} else {
			break
		}
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			}
			if s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[n1][n2]
}
