package _86

import (
    "fmt"
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

func TestPartition(t *testing.T) {
    list := makeListNode([]int{1})
    head := partition(list, 0)
    fmt.Println(toStr(head))
}
