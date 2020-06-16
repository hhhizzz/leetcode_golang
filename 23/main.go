package _23

import (
	"container/heap"
)

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

func mergeKLists1(lists []*ListNode) *ListNode {
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

type mineHeap struct {
	array []*ListNode
}

func (h *mineHeap) pop() *ListNode {
	root := 0
	last := len(h.array) - 1

	result := h.array[root]
	h.array[root], h.array[last] = h.array[last], h.array[root]
	h.array = h.array[:len(h.array)-1]

	parent := root

	for {
		child := parent*2 + 1
		if child >= len(h.array) {
			break
		}
		if rChild := child + 1; rChild < len(h.array) && h.array[rChild].Val < h.array[child].Val {
			child = rChild
		}
		if h.array[child].Val < h.array[parent].Val {
			h.array[parent], h.array[child] = h.array[child], h.array[parent]
		}
		parent = child
	}

	return result
}

func (h *mineHeap) push(node *ListNode) {
	h.array = append(h.array, node)

	child := len(h.array) - 1
	parent := (child - 1) >> 1

	for child != 0 && h.array[parent].Val > h.array[child].Val {
		h.array[child], h.array[parent] = h.array[parent], h.array[child]
		child = parent
		parent = (child - 1) >> 1
	}
}

func (h *mineHeap) nextNode() *ListNode {
	if len(h.array) == 0 {
		return nil
	}
	first := h.pop()
	next := first.Next
	if next != nil {
		h.push(next)
	}
	return first
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := mineHeap{array: []*ListNode{}}
	for _, list := range lists {
		if list != nil {
			h.push(list)
		}
	}
	front := &ListNode{}
	current := front

	for next := h.nextNode(); next != nil; next = h.nextNode() {
		current.Next = next
		current = current.Next
	}
	return front.Next
}
