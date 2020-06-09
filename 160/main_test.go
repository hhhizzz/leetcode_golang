package _60

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	headA := ListNode{1, &ListNode{-1, nil}}
	headB := ListNode{2, nil}

	intersection := ListNode{3, nil}
	headA.Next.Next = &intersection
	headB.Next = &intersection

	node := getIntersectionNode(&headA, &headB)
	fmt.Println(node.Val)
}
