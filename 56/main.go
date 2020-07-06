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
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] <= current[1] {
				current[1] = max(current[1], intervals[j][1])
				i = j
			} else {
				break
			}
		}
		res = append(res, current)
	}
	return res
}
