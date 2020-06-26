package _40

import "sort"

func helper(candidates []int, target int, current []int, res *[][]int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
		return
	}
	if len(candidates) == 0 || target < 0 {
		return
	}
	for i := 0; i < len(candidates); i++ {
		if i != 0 && candidates[i] == candidates[i-1] {
			continue
		}
		current = append(current, candidates[i])
		helper(candidates[i+1:], target-candidates[i], current, res)
		current = current[:len(current)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	helper(candidates, target, []int{}, &res)
	return res
}
