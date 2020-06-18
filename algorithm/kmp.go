package algorithm

func kmp(pattern string, text string) bool {
	if len(pattern) > len(text) {
		return false
	}
	next := make([]int, len(pattern))
	next[0] = -1
	// 自己匹配自己来构建next串，用 t = -1的位置可以匹配任意串
	for j, t := 0, -1; j < len(pattern)-1; {
		// j, t表示目前自己的匹配情况
		if t == -1 || pattern[j] == pattern[t] {
			// 如果最新的j, t 是匹配的，那么携手并进看看下一次是不是还是匹配的
			j++
			t++
			if pattern[j] == pattern[t] {
				// 如果这个下一次还是匹配的，说明pattern[:j+1] 与pattern[t-j:t+1] 是相匹配的，那么next[j]应该调到t的next位置，因为如果跳到t的位置必然也是不匹配的
				// 就多跳“一次”，表面看起来是一次，但是因为next[t]也是多跳“一次“，事实上是跳到了最远的next位置
				next[j] = next[t]
			} else {
				// 如果这个下一次没有匹配，说明pattern[:j] 与 pattern[t-j:t] 是匹配的，但是pattern[j]与pattern[t]没有匹配，这正是我们要找的next点，此时next可以跳到t的位置正好符合需求
				next[j] = t
			}
		} else {
			// 如果最新的j, t没有匹配，也就是上一次循环里发生了next[j] = t，此时也就是失配的情况，理所应当地跳转到next[t]尝试下一次匹配即可
			t = next[t]
		}
	}
	for i, j := 0, 0; i < len(text) && j < len(pattern); {
		if j == -1 || text[i] == pattern[j] {
			i++
			j++
			if j == len(pattern) {
				return true
			}
		} else {
			j = next[j]
		}

	}
	return false
}
