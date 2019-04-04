package _148

import (
    "strconv"
    "testing"
)

func makeListNode(num []int) *ListNode {
    head := ListNode{}
    last := &head
    for _, n := range num {
        node := ListNode{n, nil}
        last.Next = &node
        last = &node
    }
    return head.Next
}

func compListNode(l1 *ListNode, l2 *ListNode) bool {
    if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
        return false
    }
    for l1 != nil && l2 != nil {
        if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
            return false
        }
        if l1.Val != l2.Val {
            return false
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return true
}

func toStr(l *ListNode) string {
    result := "["
    for l != nil {
        result += strconv.Itoa(l.Val) + ", "
        l = l.Next
    }
    result += "]"
    return result
}

func TestSortList(T *testing.T) {
    grid := []*ListNode{
        makeListNode([]int{4, 2, 1, 3}),
        makeListNode([]int{-1, 5, 3, 4, 0}),
    }
    expects := []*ListNode{
        makeListNode([]int{1, 2, 3, 4}),
        makeListNode([]int{-1, 0, 3, 4, 5}),
    }
    for i, expect := range expects {
        actual := sortList(grid[i])
        if !compListNode(actual, expect) {
            T.Errorf("expected: %s,actual: %s", toStr(expect), toStr(actual))
        }
    }

}
