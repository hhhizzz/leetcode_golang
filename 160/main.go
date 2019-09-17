package _60

type ListNode struct {
    Val  int
    Next *ListNode
}

//方法是首先计算一个链表比另外一个长多少，然后长的那个先前进那个长度再同时前进看什么时候相遇
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    currentA := headA
    currentB := headB
    for currentA != nil && currentB != nil {
        if currentA == currentB {
            return currentA
        }
        currentA = currentA.Next
        currentB = currentB.Next
    }

    if currentB != nil {
        currentA, currentB = currentB, currentA
        headA, headB = headB, headA
    }
    //currentA is longer
    for currentA != nil {
        currentA = currentA.Next
        headA = headA.Next
    }
    currentB = headB
    currentA = headA
    for currentA != currentB {
        currentA = currentA.Next
        currentB = currentB.Next
    }
    return currentA
}
