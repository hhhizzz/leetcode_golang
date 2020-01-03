package _445

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    s1 := []*ListNode{}
    s2 := []*ListNode{}

    for current1 := l1; current1 != nil; current1 = current1.Next {
        s1 = append(s1, current1)
    }
    for current2 := l2; current2 != nil; current2 = current2.Next {
        s2 = append(s2, current2)
    }
    //保证l1更长
    if len(s2) > len(s1) {
        s1, s2 = s2, s1
        l1, l2 = l2, l1
    }
    carry := 0
    for len(s1) != 0 {
        current1 := s1[len(s1)-1]
        s1 = s1[:len(s1)-1]
        var current2 *ListNode
        if len(s2) != 0 {
            current2 = s2[len(s2)-1]
            s2 = s2[:len(s2)-1]
        }
        sum := current1.Val + carry
        if current2 != nil {
            sum += current2.Val
        }
        carry = sum / 10
        sum %= 10
        current1.Val = sum
    }
    if carry != 0 {
        newNode := &ListNode{Val: carry, Next: l1}
        l1 = newNode
    }
    return l1
}
