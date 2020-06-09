package _124

import (
	"math"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func back(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := back(root.Left, max)
	right := back(root.Right, max)
	current := root.Val
	if left > 0 {
		current += left
	}
	if right > 0 {
		current += right
	}
	if current > *max {
		*max = current
	}
	bigger := maxInt(left, right)
	if bigger < 0 {
		return root.Val
	} else {
		return root.Val + bigger
	}
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	back(root, &result)
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
