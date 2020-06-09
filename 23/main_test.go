package _23

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

func TestMergeKLists(t *testing.T) {
	list1 := makeListNode([]int{1, 4, 5})
	list2 := makeListNode([]int{1, 3, 4})
	list3 := makeListNode([]int{2, 6})
	result := mergeKLists([]*ListNode{list1, list2, list3})
	println(toStr(result))
}
