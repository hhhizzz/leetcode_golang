package _78

func helper(nums, current []int, n int, res *[][]int) {
	current = append(current, nums[0])
	if len(current) == n {
		temp := make([]int, n)
		copy(temp, current)
		*res = append(*res, temp)
	} else {
		for i := 1; i < len(nums); i++ {
			helper(nums[i:], current, n, res)
		}
	}
}

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			helper(nums[j:], []int{}, i, &res)
		}
	}
	return res
}
