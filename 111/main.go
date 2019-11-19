package _111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 && right == 0 {
		return 1
	} else {
		if left == 0 {
			return 1 + right
		} else if right == 0 {
			return 1 + left
		}
		return 1 + min(left, right)
	}
}
