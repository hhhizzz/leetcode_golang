package _60

type ListNode struct {
    Val  int
    Next *ListNode
}

//方法是首先计算一个链表比另外一个长多少，然后长的那个先前进那个长度再同时前进看什么时候相遇
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    length := 0
    currentA := headA
    currentB := headB
    for currentA != nil && currentB != nil {
        if currentA == currentB {
            return currentA
        }
        currentA = currentA.Next
        currentB = currentB.Next
    }
    longerList := currentA
    if currentA == nil {
        temp := headA
        headA = headB
        headB = temp
        longerList = currentB
    }
    for longerList != nil {
        longerList = longerList.Next
        length++
    }

    currentA = headA
    currentB = headB
    for length != 0 {
        currentA = currentA.Next
        length--
    }
    for currentA != currentB {
        currentA = currentA.Next
        currentB = currentB.Next
    }
    return currentA
}
