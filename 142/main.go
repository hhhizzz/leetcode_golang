package _142

type ListNode struct {
    Val  int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil {
        if fast.Next == nil {
            return nil
        }
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            newPtr := head
            for newPtr != slow {
                newPtr = newPtr.Next
                slow = slow.Next
            }
            return slow
        }
    }
    return nil
}
