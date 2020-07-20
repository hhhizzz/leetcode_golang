package _92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	length := n - m + 1
	if length == 0 {
		return head
	}
	pre := &ListNode{Next: head}
	left := pre
	for i := 1; i < m; i++ {
		left = left.Next
	}
	right1 := left
	right2 := left.Next
	for i := 0; i < length; i++ {
		right1, right2, right2.Next = right2, right2.Next, right1
	}
	left.Next, left.Next.Next = right1, right2

	return pre.Next
}
