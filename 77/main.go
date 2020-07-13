package _77

func helper(n, k int, current []int, res *[][]int) {
	if len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
	} else {
		// 这个剪枝可以从28ms -> 8ms
		diff := k - len(current)
		for i := current[len(current)-1] + 1; i+diff-1 <= n; i++ {
			helper(n, k, append(current, i), res)
		}
	}
}

func combine(n int, k int) [][]int {
	var res [][]int
	for i := 1; i <= n; i++ {
		helper(n, k, []int{i}, &res)
	}
	return res
}
