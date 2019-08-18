package _92

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n {
        return head
    }
    current := head
    var pre *ListNode
    for i := 1; i != m; i++ {
        pre = current
        current = current.Next
    }
    preM := pre
    currentM := current
    for i := m; i != n; i++ {
        next := current.Next
        current.Next = pre
        pre, current = current, next
    }
    afterN := current.Next
    currentM.Next = afterN

    current.Next = pre
    if preM != nil {
        preM.Next = current
    } else {
        head = current
    }

    return head
}
