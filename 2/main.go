package _2

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum := 0
    node1, node2 := l1, l2
    var node1Last *ListNode
    for ; node1 != nil && node2 != nil; node1, node2 = node1.Next, node2.Next {
        node1.Val = node1.Val + node2.Val + sum
        if node1.Val >= 10 {
            node1.Val -= 10
            sum = 1
        } else {
            sum = 0
        }
        node1Last = node1
    }
    if node2 != nil {
        node1Last.Next = node2
        node1 = node2
    }
    for ; node1 != nil && sum != 0; node1 = node1.Next {
        node1.Val += sum
        if node1.Val >= 10 {
            node1.Val -= 10
            sum = 1
        } else {
            sum = 0
        }
        node1Last = node1
    }
    if sum != 0 {
        node1Last.Next = &ListNode{Val: 1}
    }
    return l1
}
