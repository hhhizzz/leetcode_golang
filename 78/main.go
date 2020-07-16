package _78

import "sort"

func helper(nums []int, pos int, path []int, res *[][]int) {
	if pos == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	helper(nums, pos+1, path, res)
	helper(nums, pos+1, append(path, nums[pos]), res)
}

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	helper(nums, 0, []int{}, &res)
	return res
}
