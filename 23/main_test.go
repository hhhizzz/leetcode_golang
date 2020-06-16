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

func TestMineHeap(t *testing.T) {
	list1 := makeListNode([]int{4})
	list2 := makeListNode([]int{1})
	list3 := makeListNode([]int{2})
	h := mineHeap{array: []*ListNode{}}

	h.push(list1)
	h.push(list2)
	h.push(list3)

	var popVal int
	popVal = h.pop().Val
	if popVal != 1 {
		t.Errorf("expect %d != 1", popVal)
	}

	popVal = h.pop().Val
	if popVal != 2 {
		t.Errorf("expect %d != 2", popVal)
	}

	popVal = h.pop().Val
	if popVal != 4 {
		t.Errorf("expect %d != 4", popVal)
	}

	h.push(list1)
	popVal = h.pop().Val
	if popVal != 4 {
		t.Errorf("expect %d != 4", popVal)
	}

	h.push(list1)
	h.push(list2)
	popVal = h.pop().Val
	if popVal != 1 {
		t.Errorf("expect %d != 1", popVal)
	}

	h.push(list1)
	h.push(list2)
	h.push(list3)
	popVal = h.pop().Val
	if popVal != 1 {
		t.Errorf("expect %d != 1", popVal)
	}

}
