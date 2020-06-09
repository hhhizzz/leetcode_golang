package _543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deep(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	} else {
		lDepth := deep(root.Left, result)
		rDepth := deep(root.Right, result)
		*result = max(*result, lDepth+rDepth)
		return max(lDepth, rDepth) + 1
	}
}

//尽量避免使用全局变量，在leetcode里可能会出错
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	deep(root, &result)
	return result
}
