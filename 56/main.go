package _56

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type byStart []Interval

func (b byStart) Len() int {
	return len(b)
}

func (b byStart) Less(i, j int) bool {
	return b[i].Start < b[j].Start
}

func (b byStart) Swap(i, j int) {
	temp := b[i]
	b[i] = b[j]
	b[j] = temp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals []Interval) []Interval {
	sort.Sort(byStart(intervals))
	length := len(intervals)
	result := make([]Interval, 0) //之前的一个方案使用删除，慢了12ms
	for index := 0; index < length-1; index++ {
		if intervals[index].End >= intervals[index+1].Start {
			intervals[index+1].Start = intervals[index].Start
			intervals[index+1].End = max(intervals[index].End, intervals[index+1].End)
		} else {
			result = append(result, intervals[index])
		}
	}
	if length > 0 {
		result = append(result, intervals[length-1])
	}
	return result
}
