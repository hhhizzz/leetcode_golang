package _199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeLevel struct {
	*TreeNode
	level int
}

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := []*TreeNodeLevel{&TreeNodeLevel{TreeNode: root, level: 0}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		if len(queue) == 0 || queue[0].level != current.level {
			result = append(result, current.Val)
		}
		if current.Left != nil {
			queue = append(queue, &TreeNodeLevel{TreeNode: current.Left, level: current.level + 1})
		}
		if current.Right != nil {
			queue = append(queue, &TreeNodeLevel{TreeNode: current.Right, level: current.level + 1})
		}
	}
	return result
}
