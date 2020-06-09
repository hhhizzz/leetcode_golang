package _2

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

//新加了一个prev节点用于保持l1节点是更长的
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var prev1 *ListNode
	node1 := l1
	node2 := l2

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	carry := 0
	for node1 != nil && node2 != nil {
		newVal := node1.Val + node2.Val + carry
		carry = newVal / 10
		newVal = newVal % 10

		node1.Val = newVal

		prev1 = node1
		node1 = node1.Next
		node2 = node2.Next
	}

	//这里有一个警告，说明golang还不够智能，如果能达到这里上一个循环至少会执行一次这个prev必然不可能为nil
	if node1 == nil {
		prev1.Next = node2
		node1 = node2
	}

	for node1 != nil {
		newVal := node1.Val + carry
		carry = newVal / 10
		newVal = newVal % 10

		node1.Val = newVal

		prev1 = node1
		node1 = node1.Next
	}
	if carry != 0 && prev1 != nil {
		prev1.Next = &ListNode{Val: carry}
	}

	return l1
}
