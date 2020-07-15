package _86

import (
	"strconv"
)

func toStr(l *ListNode) string {
	result := "["
	for l != nil {
		result += strconv.Itoa(l.Val) + ", "
		l = l.Next
	}
	result += "]"
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	leftPre := &ListNode{}
	rightPre := &ListNode{}

	left := leftPre
	right := rightPre
	for current := head; current != nil; {
		if current.Val < x {
			left.Next = current
			left = left.Next
		} else {
			right.Next = current
			right = right.Next
		}
		current, current.Next = current.Next, nil
	}
	left.Next = rightPre.Next
	return leftPre.Next
}
