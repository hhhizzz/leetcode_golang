package _133

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	stack := []*Node{node}
	myNodes := map[int]*Node{node.Val: &Node{Val: node.Val, Neighbors: []*Node{}}}
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		newNode := myNodes[current.Val]
		for _, nextNode := range current.Neighbors {
			if NewNextNode, ok := myNodes[nextNode.Val]; ok {
				newNode.Neighbors = append(newNode.Neighbors, NewNextNode)
			} else {
				NewNextNode := &Node{Val: nextNode.Val, Neighbors: []*Node{}}
				newNode.Neighbors = append(newNode.Neighbors, NewNextNode)
				stack = append(stack, nextNode)
				myNodes[NewNextNode.Val] = NewNextNode
			}
		}
	}
	return myNodes[node.Val]
}
