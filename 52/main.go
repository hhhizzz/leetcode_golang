package _52

func dfs(row, col, slash, bSlash, n uint32, res *int) {
	// 输入数据 0 表示可用状态
	if row >= n {
		*res += 1
		return
	}
	var bitsN uint32 = (1 << n) - 1
	// bits中1表示可用状态
	bits := ^(col | slash | bSlash) & bitsN
	for bits != 0 {
		// 取出最低一位的1
		lowBit := bits & -bits
		// 每下降一行，slash往左一格，bSlash往右一格
		dfs(row+1, col|lowBit, (slash|lowBit)<<1, (bSlash|lowBit)>>1, n, res)
		// 去掉这个最低位，尝试下一个
		bits -= lowBit
	}
}

func totalNQueens(n int) int {
	res := 0
	dfs(0, 0, 0, 0, uint32(n), &res)
	return res
}
