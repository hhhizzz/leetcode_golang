package _3

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	// m 表示某个字符上次出现的位置
	m := map[byte]int{}
	result := 0
	// l 到 r 表示当前符合要求的字符串，因此当前长度为 l - r + 1
	for l, r := 0, 0; r < len(s); r++ {
		c := s[r]
		// 检查当前字符串是否符合要求，方法是检查当前c上次出现的位置是不是在l右边
		if lastR, ok := m[c]; ok {
			l = max(l, lastR+1)
		}
		// 通过当前字符串的长度更新结果
		result = max(result, r-l+1)
		m[c] = r
	}
	return result
}
