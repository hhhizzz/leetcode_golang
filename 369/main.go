package _369

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    prev, current := head, head.Next
    for current != nil {
        next := current.Next
        current.Next = prev

        prev, current = current, next
    }
    head.Next = nil
    return prev
}

func plusOne(head *ListNode) *ListNode {
    head = reverse(head)
    current := head

    carry := 1
    for current.Next != nil {
        current.Val += carry
        carry = current.Val / 10
        current.Val %= 10
        current = current.Next
    }
    current.Val += carry
    carry = current.Val / 10
    current.Val %= 10
    if carry != 0 {
        current.Next = &ListNode{Val: carry}
    }
    return reverse(head)
}
