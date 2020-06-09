package _148

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	if length < 2 {
		return head
	}

	for step := 1; step < length; step *= 2 {
		current = head
		pre := &ListNode{Next: current}
		for current != nil {
			left := current
			right := cut(current, step)
			current = cut(right, step)
			next := merge(left, right)
			if pre.Next == head {
				pre.Next = next
				head = next
			} else {
				pre.Next = next
			}
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return head
}

func merge(a *ListNode, b *ListNode) *ListNode {
	current := &ListNode{}
	result := current
	for a != nil && b != nil {
		if a.Val < b.Val {
			current.Next = a
			a = a.Next
		} else {
			current.Next = b
			b = b.Next
		}
		current = current.Next
	}
	if a != nil {
		current.Next = a
	} else {
		current.Next = b
	}
	return result.Next
}

func cut(l *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		if l != nil {
			if i == n-1 {
				temp := l.Next
				l.Next = nil
				l = temp
			} else {
				l = l.Next
			}

		}
	}
	return l
}
