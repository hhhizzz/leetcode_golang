package _253

import (
	"container/heap"
	"sort"
)

type roomHeap []int

func (h roomHeap) Len() int {
	return len(h)
}

func (h roomHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h roomHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *roomHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *roomHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})

	h := &roomHeap{intervals[0][1]}
	heap.Init(h)

	result := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, intervals[i][1])
		if len(*h) > result {
			result = len(*h)
		}
	}
	return result
}
