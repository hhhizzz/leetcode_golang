package _90

import "sort"

// 函数表示从pos开始枚举查找子集，path表示已经枚举过的路径
func helper(nums []int, pos int, path []int, res *[][]int) {
	// 递归尾部 把当前path放入结果
	if pos == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	// nums[back]表示右边与nums[pos]不同的第一个数，也就是nums[pos:back]都是相同的数
	back := pos
	for back < len(nums) && nums[back] == nums[pos] {
		back++
	}
	// 枚举如果pos没在path里的情况
	helper(nums, back, path, res)
	// 枚举path里分别出现多次的情况，这里是与nums没有重复不一样的地方
	// 因为nums有重复，因此在递归的时候直接跳过重复数，单独把重复数字放到path里
	for i := pos; i < back; i++ {
		path = append(path, nums[i])
		helper(nums, back, path, res)
	}
	path = path[:len(path)-back+pos]
}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	helper(nums, 0, []int{}, &res)
	return res
}
