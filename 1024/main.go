package _1024

import "sort"

//贪心法，从0开始，每次找能到达的最长的视频片段即可
func videoStitching(clips [][]int, T int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})
	current := 0
	result := 0
	i := 0
	next := 0
	for current < T {
		for ; i < len(clips) && clips[i][0] <= current; i++ {
			if clips[i][1] > next {
				next = clips[i][1]
			}
		}
		if current == next {
			return -1
		}
		current = next
		result++
		if i == len(clips) && current < T {
			return -1
		}
	}
	return result
}
