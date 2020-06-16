package _24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	front := &ListNode{Next: head}
	pre := front
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := first.Next
		// pre -> first -> second ====> pre -> second -> first
		pre.Next, first.Next, second.Next = second, second.Next, first
		pre = first
	}
	return front.Next
}
