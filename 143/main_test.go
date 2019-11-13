package _143

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

func TestFindMid(t *testing.T) {
    var list *ListNode
    var mid *ListNode

    list = makeListNode([]int{1})
    mid = findMid(list)
    if mid.Val != 1 {
        t.Errorf("expected %d,actual %d in %v\n", 1, mid.Val, toStr(list))
    }

    list = makeListNode([]int{1, 2})
    mid = findMid(list)
    if mid.Val != 2 {
        t.Errorf("expected %d,actual %d in %v\n", 2, mid.Val, toStr(list))
    }

    list = makeListNode([]int{1, 2, 3})
    mid = findMid(list)
    if mid.Val != 3 {
        t.Errorf("expected %d,actual %d in %v\n", 3, mid.Val, toStr(list))
    }
    list = makeListNode([]int{1, 2, 3, 4})
    mid = findMid(list)
    if mid.Val != 3 {
        t.Errorf("expected %d,actual %d in %v\n", 3, mid.Val, toStr(list))
    }
    list = makeListNode([]int{1, 2, 3, 4, 5})
    mid = findMid(list)
    if mid.Val != 4 {
        t.Errorf("expected %d,actual %d in %v\n", 5, mid.Val, toStr(list))
    }
}

func TestReverse(t *testing.T) {
    var list *ListNode

    list = makeListNode([]int{1})
    fmt.Println(toStr(reverseList(list)))

    list = makeListNode([]int{1, 2})
    fmt.Println(toStr(reverseList(list)))

    list = makeListNode([]int{1, 2, 3})
    fmt.Println(toStr(reverseList(list)))

    list = makeListNode([]int{1, 2, 3, 4})
    fmt.Println(toStr(reverseList(list)))

    list = makeListNode([]int{1, 2, 3, 4, 5})
    fmt.Println(toStr(reverseList(list)))
}

func TestReorderList(t *testing.T) {
    var list *ListNode

    list = makeListNode([]int{1})
    reorderList(list)
    fmt.Println(toStr(list))

    list = makeListNode([]int{1, 2})
    reorderList(list)
    fmt.Println(toStr(list))

    list = makeListNode([]int{1, 2, 3})
    reorderList(list)
    fmt.Println(toStr(list))

    list = makeListNode([]int{1, 2, 3, 4})
    reorderList(list)
    fmt.Println(toStr(list))

    list = makeListNode([]int{1, 2, 3, 4, 5})
    reorderList(list)
    fmt.Println(toStr(list))

}
