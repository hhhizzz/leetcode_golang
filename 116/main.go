package _116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	level := root
	for level.Left != nil {
		current := level
		for current != nil {
			current.Left.Next = current.Right
			if current.Next != nil {
				current.Right.Next = current.Next.Left
			}
			current = current.Next
		}
		level = level.Left
	}
	return root
}
