package _132

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//动态规划
func minCut(s string) int {
	//dp[i][j]表示s[i]到s[j]是不是回文
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for length := 0; length < len(s); length++ {
		for left := 0; left < len(s); left++ {
			right := left + length
			if right >= len(s) {
				break
			}
			if s[left] == s[right] {
				if length < 3 || dp[left+1][right-1] {
					dp[left][right] = true
				}
			}
		}
	}
	//cut[i]表示s[i]到s[len(s)-1]的最小切割数
	cut := make([]int, len(s))
	for i := range cut {
		cut[i] = len(s) - 1 - i
	}
	for start := len(s) - 1; start >= 0; start-- {
		for end := start; end < len(s); end++ {
			if dp[start][end] {
				if end == len(s)-1 {
					cut[start] = 0
				} else {
					cut[start] = min(cut[start], cut[end+1]+1)
				}
			}
		}
	}
	return cut[0]
}

//中间展开法
func minCut2(s string) int {
	//dp[i]表示s[0]到s[i]的最小切割次数，目标求dp[len(s)-1]
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = i
	}
	for i := 0; i < len(s); i++ {
		length := 0
		//从中间展开
		for {
			left := i - length
			right := i + length
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] == s[right] {
				if left == 0 {
					dp[right] = 0
				} else {
					dp[right] = min(dp[right], dp[left-1]+1)
				}
			} else {
				break
			}
			length++
		}
		length = 0
		//和右边一起作为中间展开，不和左边是因为上一次循环已经和左边一起展开过了
		for {
			left := i - length
			right := i + length + 1
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] == s[right] {
				if left == 0 {
					dp[right] = 0
				} else {
					dp[right] = min(dp[right], dp[left-1]+1)
				}
			} else {
				break
			}
			length++
		}
	}
	return dp[len(s)-1]
}
