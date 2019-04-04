package _148

type ListNode struct {
    Val  int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    preSlow, slow, fast := head, head, head
    for fast != nil && fast.Next != nil {
        preSlow = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    preSlow.Next = nil
    return merge(sortList(head), sortList(slow))
}
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val < l2.Val {
        l1.Next = merge(l1.Next, l2)
        return l1
    } else {
        l2.Next = merge(l2.Next, l1)
        return l2
    }
}
