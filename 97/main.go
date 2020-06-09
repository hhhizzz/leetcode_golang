package _97

func helper(s1, s2, s3 []byte) bool {
	if len(s3) == 0 {
		return true
	}
	var result bool
	if len(s1) > 0 && s3[0] == s1[0] {
		result = helper(s1[1:], s2, s3[1:])
	}
	if result {
		return true
	} else {
		if len(s2) > 0 && s3[0] == s2[0] {
			return helper(s1, s2[1:], s3[1:])
		} else {
			return false
		}
	}
}

//方法一：直接深搜
func isInterleave1(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return helper([]byte(s1), []byte(s2), []byte(s3))
}

//方法二：动态规划
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	array1 := []byte(s1)
	array2 := []byte(s2)
	array3 := []byte(s3)

	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && array1[i-1] == array3[i-1]
	}
	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && array2[j-1] == array3[j-1]
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i][j] = (dp[i-1][j] && array1[i-1] == array3[i+j-1]) || (dp[i][j-1] && array2[j-1] == array3[i+j-1])
		}
	}
	//fmt.Println(dp)
	return dp[len(s1)][len(s2)]
}
