package _206

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    //注意pre应该是nil
    var pre *ListNode
    current := head

    for next := head.Next; current.Next != nil; next = current.Next {
        current.Next = pre
        pre = current
        current = next
    }
    current.Next = pre
    pre = current
    return pre
}
