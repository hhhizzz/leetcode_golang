package _138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	current := head
	newHead := &Node{Val: head.Val}
	newCurrent := newHead
	m := map[*Node]*Node{current: newCurrent}
	for current.Next != nil {
		if newRandom, ok := m[current.Random]; !ok && current.Random != nil {
			newCurrent.Random = &Node{Val: current.Random.Val}
			m[current.Random] = newCurrent.Random
		} else {
			newCurrent.Random = newRandom
		}
		if newNext, ok := m[current.Next]; !ok {
			newCurrent.Next = &Node{Val: current.Next.Val}
			m[current.Next] = newCurrent.Next
		} else {
			newCurrent.Next = newNext
		}
		newCurrent = newCurrent.Next
		current = current.Next
	}
	if current.Random != nil {
		newCurrent.Random = m[current.Random]
	}
	return newHead
}
