package _328

import "testing"

func TestOddEvenList(t *testing.T) {
	a := &ListNode{1, nil}
	a.Next = &ListNode{2, nil}
	a.Next.Next = &ListNode{3, nil}
	oddEvenList(a)
}
