package _206

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var pre *ListNode
    current, next := head, head.Next
    for current != nil {
        current.Next = pre
        pre = current
        current = next
        if next!=nil{
            next = next.Next
        }
    }
    return pre
}
