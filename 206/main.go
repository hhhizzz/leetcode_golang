package _206

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    var pre *ListNode
    current := head
    var next *ListNode

    for current != nil{
        next = current.Next
        current.Next = pre
        pre = current
        current = next
    }
    return pre
}
