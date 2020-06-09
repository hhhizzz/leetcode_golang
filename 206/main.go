package _206

type ListNode struct {
	Val  int
	Next *ListNode
}

//非递归版本
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var current *ListNode = head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev, current = current, next
	}
	return prev
}

//递归版本
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
