package _295

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

type MedianFinder struct {
	maxHeap
	minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	h1 := maxHeap{array: []int{}}
	heap.Init(&h1)
	h2 := minHeap{array: []int{}}
	heap.Init(&h2)
	return MedianFinder{h1, h2}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.maxHeap, num)
	if this.maxHeap.Len() > this.minHeap.Len() {
		median := heap.Pop(&this.maxHeap)
		heap.Push(&this.minHeap, median)
	}
	if this.minHeap.Len() > this.maxHeap.Len() {
		median := heap.Pop(&this.minHeap)
		heap.Push(&this.maxHeap, median)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64(this.minHeap.array[0]+this.maxHeap.array[0]) / 2
	} else {
		return float64(this.maxHeap.array[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
