package _91

func numDecodings(s string) int {
	dp := make([]int, len(s))
	if s[0] > '0' {
		dp[0] = 1
	} else {
		return 0
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] == '1' {
			if s[i] > '0' {
				if i-2 >= 0 {
					dp[i] = dp[i-1] + dp[i-2]
				} else {
					dp[i] = dp[i-1] + 1
				}
			} else {
				if i-2 >= 0 {
					dp[i] = dp[i-2]
				} else {
					dp[i] = 1
				}
			}
		} else if s[i-1] == '2' {
			if s[i] <= '6' && s[i] > '0' {
				if i-2 >= 0 {
					dp[i] = dp[i-1] + dp[i-2]
				} else {
					dp[i] = dp[i-1] + 1
				}
			} else if s[i] == '0' {
				if i-2 >= 0 {
					dp[i] = dp[i-2]
				} else {
					dp[i] = 1
				}
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			if s[i] == '0' {
				return 0
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[len(s)-1]
}
