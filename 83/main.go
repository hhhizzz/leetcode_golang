package _83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		if current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
