package _92

import (
    "fmt"
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

func toStr(l *ListNode) string {
    result := "["
    for l != nil {
        result += strconv.Itoa(l.Val) + ", "
        l = l.Next
    }
    result += "]"
    return result
}

func TestReverse(t *testing.T) {
    l:=reverseBetween(makeListNode([]int{1, 2, 3, 4, 5}), 5, 5)
    fmt.Println(toStr(l))
}
