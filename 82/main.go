package _82

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	pre := &ListNode{}
	last := pre
	for current != nil {
		if current.Next != nil && current.Next.Val == current.Val {
			val := current.Val
			for current != nil && current.Val == val {
				current = current.Next
			}
		} else {
			last.Next = current
			last = current
			current = current.Next
			last.Next = nil
		}
	}
	return pre.Next
}
