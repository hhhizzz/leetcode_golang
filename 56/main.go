package _56

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		current := []int{intervals[i][0], intervals[i][1]}
		if len(res) == 0 || intervals[i][0] > res[len(res)-1][1] {
			res = append(res, current)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}
