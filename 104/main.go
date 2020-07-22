package _104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeLevel struct {
	Level int
	Node  *TreeNode
}

func maxDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}

	var stack []NodeLevel
	stack = append(stack, NodeLevel{Level: 1, Node: root})
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Level > result {
			result = current.Level
		}
		if current.Node.Left != nil {
			stack = append(stack, NodeLevel{Level: current.Level + 1, Node: current.Node.Left})
		}
		if current.Node.Right != nil {
			stack = append(stack, NodeLevel{Level: current.Level + 1, Node: current.Node.Right})
		}
	}
	return result
}
