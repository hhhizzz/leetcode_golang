package _78

func helper(nums, current []int, pos, n int, res *[][]int) {
	current = append(current, nums[pos])
	if len(current) == n {
		temp := make([]int, n)
		copy(temp, current)
		*res = append(*res, temp)
	} else {
		diff := n - len(current)
		for i := pos + 1; i+diff-1 < len(nums); i++ {
			helper(nums, current, i, n, res)
		}
	}
}

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			helper(nums, []int{}, j, i, &res)
		}
	}
	return res
}
