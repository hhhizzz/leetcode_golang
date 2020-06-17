package _25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	front := &ListNode{Next: head}

	pre, current := front, front.Next
	for current != nil {
		tail := pre
		for i := 0; tail != nil && i < k; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		back := tail.Next
		next := current.Next
		for i := 0; i < k-1; i++ {
			temp := next.Next
			next.Next = current
			current = next
			next = temp
		}
		pre, pre.Next, pre.Next.Next = pre.Next, tail, back
		current = back
	}
	return front.Next
}
