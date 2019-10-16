package _109

import (
    "fmt"
    "testing"
)

func TestSortedListToBST(t *testing.T) {
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    root := sortedListToBST(head)
    fmt.Println(root)
}
