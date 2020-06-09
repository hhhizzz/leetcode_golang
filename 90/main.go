package _90

import "sort"

func dfs(nums []int, left int, current []int, result *[][]int) {
	if left > len(nums) {
		return
	}
	if left == 0 {
		newArray := make([]int, len(current))
		copy(newArray, current)
		*result = append(*result, newArray)
	} else {
		for i := 0; i < len(nums); i++ {
			if i != 0 && nums[i-1] == nums[i] {
				continue
			}
			current = append(current, nums[i])
			dfs(nums[i+1:], left-1, current, result)
			current = current[:len(current)-1]
		}
	}
}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	sort.Ints(nums)
	for length := 1; length <= len(nums); length++ {
		dfs(nums, length, []int{}, &result)
	}
	return result
}
