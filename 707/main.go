package _707

type node struct {
    Val  int
    Prev *node
    Next *node
}

type MyLinkedList struct {
    Front  *node
    Back   *node
    Length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    myLinkedList := MyLinkedList{Front: &node{Val: 0}, Back: &node{Val: 0}, Length: 0}
    myLinkedList.Front.Next = myLinkedList.Back
    myLinkedList.Back.Prev = myLinkedList.Front
    return myLinkedList
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.Length {
        return -1
    }
    current := this.Front
    for i := 0; i <= index; i++ {
        current = current.Next
    }
    return current.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
    oldHead := this.Front.Next
    newHead := &node{Val: val}

    this.Front.Next = newHead

    newHead.Prev = this.Front
    newHead.Next = oldHead

    oldHead.Prev = newHead

    this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
    oldTail := this.Back.Prev
    newTail := &node{Val: val}

    oldTail.Next = newTail

    newTail.Prev = oldTail
    newTail.Next = this.Back

    this.Back.Prev = newTail

    this.Length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index <= 0 {
        this.AddAtHead(val)
        return
    }
    if index > this.Length {
        return
    }
    current := this.Front
    for i := 0; i < index; i++ {
        current = current.Next
    }
    next := current.Next
    newNode := &node{Val: val}

    current.Next = newNode
    newNode.Prev = current
    newNode.Next = next
    next.Prev = newNode

    this.Length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.Length {
        return
    }
    current := this.Front
    for i := 0; i < index; i++ {
        current = current.Next
    }
    next := current.Next
    nextNext := next.Next

    current.Next = nextNext
    nextNext.Prev = current

    this.Length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
