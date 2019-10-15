package _86

import (
    "strconv"
)

func toStr(l *ListNode) string {
    result := "["
    for l != nil {
        result += strconv.Itoa(l.Val) + ", "
        l = l.Next
    }
    result += "]"
    return result
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }
    current := head

    pre := &ListNode{Next: head}
    current = head

    tail := &ListNode{}
    currentTail := tail

    for current != nil {
        if current.Val >= x {
            if current == head {
                head = head.Next
            }
            pre.Next = current.Next
            currentTail.Next = current
            currentTail = current
            current.Next = nil
            current = pre.Next
        } else {
            pre, current = pre.Next, current.Next
        }
    }
    if head == nil {
        head = tail.Next
    } else {
        pre.Next = tail.Next
    }
    return head
}
