package _703

import "container/heap"

type heapInt []int

// Len is the number of elements in the collection.
func (h heapInt) Len() int {
	return len(h)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (h heapInt) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h heapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapInt) Top() int {
	x := (*h)[0]
	return x
}

func (h *heapInt) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return x
}

type KthLargest struct {
	heapInt
	k int
}

func Constructor(k int, nums []int) KthLargest {
	this := KthLargest{[]int{}, k}
	heap.Init(&this)
	for _, num := range nums {
		this.Add(num)
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	if this.Len() < this.k {
		heap.Push(this, val)
		return this.Top()
	}
	min := this.Top()
	if val > min {
		heap.Pop(this)
		heap.Push(this, val)
		min = this.Top()
	}
	return min
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
