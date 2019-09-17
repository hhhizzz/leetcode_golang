package _328

type ListNode struct {
    Val  int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
    oddList := &ListNode{Val: 0, Next: nil}
    evenList := &ListNode{Val: 0, Next: nil}

    var oddCurrent = oddList
    var evenCurrent = evenList

    current := head

    odd := true
    for current != nil {
        if odd {
            oddCurrent.Next = current
            oddCurrent = current
        } else {
            evenCurrent.Next = current
            evenCurrent = current
        }
        odd = !odd
        current = current.Next
    }
    oddCurrent.Next = evenList.Next
    evenCurrent.Next = nil
    return oddList.Next
}
