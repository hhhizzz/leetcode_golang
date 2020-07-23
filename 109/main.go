package _109

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	pre := &ListNode{Next: head}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = pre.Next
		slow = slow.Next
	}
	root := &TreeNode{Val: slow.Val}
	pre.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)

	return root
}
