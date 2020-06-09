package _110

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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

func height(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	} else {
		var balanced bool
		leftHeight, balanced := height(root.Left)
		if !balanced {
			return -1, false
		}
		rightHeight, balanced := height(root.Right)
		if !balanced {
			return -1, false
		}
		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			return -1, false
		}
		return max(leftHeight, rightHeight) + 1, true
	}
}

func isBalanced(root *TreeNode) bool {
	_, balanced := height(root)
	return balanced
}
