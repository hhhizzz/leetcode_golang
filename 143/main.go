package _143

type ListNode struct {
	Val  int
	Next *ListNode
}

func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		if fast == nil {
			//偶数长度
			preSlow := slow
			slow = slow.Next
			preSlow.Next = nil
		} else if fast.Next == nil {
			//奇数长度
			slow = slow.Next
			preSlow := slow
			slow = slow.Next
			preSlow.Next = nil
		} else {
			slow = slow.Next
		}
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	current := head
	for current != nil {
		next := current.Next
		newHead.Next, current.Next = current, newHead.Next
		current = next
	}
	return newHead.Next
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := findMid(head)
	newMid := reverseList(mid)

	current := head
	for newMid != nil {
		next := newMid.Next
		current.Next, newMid.Next = newMid, current.Next
		current = newMid.Next
		newMid = next
	}
}
