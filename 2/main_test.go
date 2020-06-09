package _2

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

func toStr(l *ListNode) string {
	result := "["
	for l != nil {
		result += strconv.Itoa(l.Val) + ", "
		l = l.Next
	}
	result += "]"
	return result
}

func TestAddTwoList(t *testing.T) {
	l := addTwoNumbers(makeListNode([]int{1}), makeListNode([]int{9, 9}))
	println(toStr(l))
}
