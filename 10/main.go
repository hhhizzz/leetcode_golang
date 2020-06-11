package _10

func match(sByte, pByte byte) bool {
	return sByte == pByte || pByte == '.'
}

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	f := make([][]bool, n+1)
	for i := 0; i <= len(s); i++ {
		f[i] = make([]bool, m+1)
	}

	//空字符串必然匹配空字符串
	f[0][0] = true
	s = " " + s
	p = " " + p

	for i := 0; i < len(s); i++ {
		// j从1开始尝试匹配，因为空字符串必然无法匹配非空字符串，也就是f[0][1:m+1]必然都是false
		for j := 1; j < len(p); j++ {
			//针对i==0的情况单独处理
			if i == 0 {
				//如果出现*，那么检查前面是否匹配，例如a*b*,这里如果a*能匹配上，那么b*一定也能匹配上，反之亦然
				if p[j] == '*' {
					f[i][j] = f[i][j-2]
				}
			} else {
				// 如果当前p串不是*，平凡的状态转移，也就是当前能匹配上以及p[j-1]能匹配s[i-1]，当前就能匹配
				if p[j] != '*' {
					f[i][j] = f[i-1][j-1] && match(s[i], p[j])
				} else {
					// 较为复杂一点的状态转移，分为两种情况
					//   1. 匹配零个，也就是和上面i==0的情况一样的处理
					//   2. 匹配多个，这是本题最为复杂的一个地方，需要检查两个地方:
					//     2.1. `*`前面一个字符也就是p[j-1]是否匹配当前s[i]
					//     2.2. s往前的部分(s[0:i])是否能匹配当前的p[0:j+1]，也就是求 f[i-1][j]，这个值上一次循环已经求过，因此可以直接使用
					// 出现了`*` j一定是大于1的，否则p串不合法，因此这里不需要判断j-2的范围
					f[i][j] = f[i][j-2] || (match(s[i], p[j-1]) && f[i-1][j])
				}
			}

		}
	}
	return f[n][m]
}
