package _117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	level := root
	for level != nil {
		current := level
		nextLevelPre := &Node{}
		nextLevelCurrent := nextLevelPre
		for current != nil {
			if current.Left != nil {
				nextLevelCurrent.Next = current.Left
				nextLevelCurrent = nextLevelCurrent.Next
			}
			if current.Right != nil {
				nextLevelCurrent.Next = current.Right
				nextLevelCurrent = nextLevelCurrent.Next
			}
			current = current.Next
		}
		level = nextLevelPre.Next
	}
	return root
}
