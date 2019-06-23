package _24

type ListNode struct {
    Val  int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    if head.Next==nil{
        return head
    }
    var pre = &ListNode{Next: head}
    current := head
    next := current.Next
    headNext := next
    for current != nil && next != nil {
        current.Next = next.Next
        next.Next = current
        pre.Next = next
        pre = current

        current = pre.Next
        if current != nil {
            next = current.Next
        }
    }
    return headNext
}
