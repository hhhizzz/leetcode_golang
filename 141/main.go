package _141

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    slow := head
    fast := head
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
        if fast == nil {
            return false
        }
        fast = fast.Next
        if slow == fast {
            return true
        }
    }
    return false
}
