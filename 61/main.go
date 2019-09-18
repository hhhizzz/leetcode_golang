package _61

type ListNode struct {
    Val  int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    var current = head
    length := 0
    for current.Next != nil {
        length++
        current = current.Next
    }
    current.Next = head
    k %= length
    k = length-k
    for i := 0; i < k; i++ {
        head = head.Next
        current = current.Next
    }
    current.Next = nil
    return head
}
