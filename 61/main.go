package _61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	current := head
	length := 1
	for current.Next != nil {
		length++
		current = current.Next
	}
	k %= length
	k = length - k
	current.Next = head
	for i := 0; i < k; i++ {
		head = head.Next
		current = current.Next
	}
	current.Next = nil
	return head
}
