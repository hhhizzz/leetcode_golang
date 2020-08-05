package _129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, result *int, current int) {
	if root == nil {
		return
	}
	current *= 10
	current += root.Val
	if root.Left == nil && root.Right == nil {
		*result += current
	} else {
		helper(root.Left, result, current)
		helper(root.Right, result, current)
	}
}

func sumNumbers(root *TreeNode) int {
	var result int
	helper(root, &result, 0)
	return result
}
