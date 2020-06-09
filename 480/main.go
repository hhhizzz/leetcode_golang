package _480

import "container/heap"

type maxHeap struct {
	array []int
}

func (h *maxHeap) Push(x interface{}) {
	h.array = append(h.array, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	result := h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	return result
}

func (h *maxHeap) Len() int {
	return len(h.array)
}

func (h *maxHeap) Less(i, j int) bool {
	return h.array[i] > h.array[j]
}

func (h *maxHeap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

type minHeap struct {
	array []int
}

func (h *minHeap) Push(x interface{}) {
	h.array = append(h.array, x.(int))
}

func (h *minHeap) Pop() interface{} {
	result := h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	return result
}

func (h *minHeap) Len() int {
	return len(h.array)
}

func (h *minHeap) Less(i, j int) bool {
	return h.array[i] < h.array[j]
}

func (h *minHeap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func medianSlidingWindow(nums []int, k int) []float64 {
	maxH := maxHeap{array: []int{}}
	minH := minHeap{array: []int{}}

	var result []float64

	for i := 0; i < len(nums); i++ {
		heap.Push(&maxH, nums[i])
		if maxH.Len() > minH.Len() {
			median := heap.Pop(&maxH)
			heap.Push(&minH, median)
		}
		if minH.Len() > maxH.Len() {
			median := heap.Pop(&minH)
			heap.Push(&maxH, median)
		}
		if i >= k-1 {
			var median float64
			if k%2 == 0 {
				median = float64(maxH.array[0]+minH.array[0]) / 2
			} else {
				median = float64(maxH.array[0])
			}
			result = append(result, median)
			if nums[i-k+1] <= maxH.array[0] {
				for j := 0; j < len(maxH.array); j++ {
					if nums[i-k+1] == maxH.array[j] {
						heap.Remove(&maxH, j)
						if minH.Len() > maxH.Len() {
							median := heap.Pop(&minH)
							heap.Push(&maxH, median)
						}
						break
					}
				}
			} else {
				for j := 0; j < len(minH.array); j++ {
					if nums[i-k+1] == minH.array[j] {
						heap.Remove(&minH, j)
						if maxH.Len() > minH.Len()+1 {
							median := heap.Pop(&maxH)
							heap.Push(&minH, median)
						}
						break
					}
				}
			}
		}
	}
	return result
}
