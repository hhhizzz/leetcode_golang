package _382

import "math/rand"

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

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	result := this.head.Val
	current := this.head.Next
	for k := 2; current != nil; k++ {
		chance := 1 / float64(k)
		happen := rand.Float64()
		if happen < chance {
			result = current.Val
		}
		current = current.Next
	}
	return result
}
