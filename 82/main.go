package _82

type ListNode struct {
	Val  int
	Next *ListNode
}

//注意head如果被删除需要移动head
func deleteDuplicates(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			val := current.Val
			for current != nil && current.Val == val {
				current = current.Next
			}
			if head.Val == val {
				head = current
			}
			pre.Next = current
		} else {
			pre = pre.Next
			current = pre.Next
		}
	}
	return head
}
