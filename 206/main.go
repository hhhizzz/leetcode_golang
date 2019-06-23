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
    current, next := head, head.Next
    for current != nil {
        current.Next = pre
        pre = current
        current = next
        if current != nil {
            next = current.Next
        }
    }
    return pre
}
