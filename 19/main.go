package _19

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fast := head
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    if fast==nil{
        return head.Next
    }
    slow := head
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    slow.Next = slow.Next.Next
    return head
}
