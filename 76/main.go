package _76

// 双指针法，相比直接做了一点优化
func minWindow(s string, t string) string {
	// 使用一个map来存储字符出现的次数，这里的优化是只用这一个变量，在计数的时候用减的方式，下面cnt同理
	m := map[byte]int{}
	// cnt表示字符的个数
	cnt := 0
	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; !ok {
			cnt++
		}
		m[t[i]]++
	}
	// l r 表示左右区间
	l := 0
	r := 0
	var res string
	for r < len(s) {
		// 遍历的时候对每一个字符出现次数减一，如果减到0说明与t中的出现次数相同了，此时可以减cnt了
		m[s[r]] -= 1
		if m[s[r]] == 0 {
			cnt--
		}
		// 检查l是否可以滑动，注意只有m[s[l]]为负的时候才能滑动，也就是要么没在t里，要么出现的次数比t里多了
		for l < r && m[s[l]] < 0 {
			m[s[l]]++
			l++
		}
		// 如果cnt减到零也就是达到目标，尝试更新res
		if cnt == 0 {
			if len(res) == 0 || r-l+1 < len(res) {
				res = s[l : r+1]
			}
		}
		r++
	}
	return res
}
