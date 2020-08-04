package _124

import (
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(root *TreeNode, result *int) int {
	current := root.Val
	var leftPath int
	if root.Left != nil {
		leftPath = helper(root.Left, result)
		if leftPath > 0 {
			current += leftPath
		}
	}
	var rightPath int
	if root.Right != nil {
		rightPath = helper(root.Right, result)
		if rightPath > 0 {
			current += rightPath
		}
	}
	*result = max(*result, current)
	return max(root.Val, root.Val+max(leftPath, rightPath))
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	helper(root, &result)
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
