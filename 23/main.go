package _23

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type nodeHeap []*ListNode

func (h nodeHeap) Len() int {
	return len(h)
}
func (h nodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h nodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	heap.Init(h)
	front := &ListNode{}
	current := front
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	for h.Len() != 0 {
		min := (*h)[0]
		current.Next = min
		heap.Pop(h)
		if min.Next != nil {
			heap.Push(h, min.Next)
		}
		current = current.Next
	}
	return front.Next
}
