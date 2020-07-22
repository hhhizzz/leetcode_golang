package _102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeLevel struct {
	Level int
	Node  *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []NodeLevel{{Level: 0, Node: root}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for len(result) <= current.Level {
			result = append(result, []int{})
		}
		result[current.Level] = append(result[current.Level], current.Node.Val)
		if current.Node.Left != nil {
			queue = append(queue, NodeLevel{Level: current.Level + 1, Node: current.Node.Left})
		}
		if current.Node.Right != nil {
			queue = append(queue, NodeLevel{Level: current.Level + 1, Node: current.Node.Right})
		}
	}
	return result
}
