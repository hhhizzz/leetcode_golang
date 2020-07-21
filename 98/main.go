package _98

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

func isValid(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}
	return isValid(root.Left, nil, &root.Val) && isValid(root.Right, &root.Val, nil)
}

func isValidBST1(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

// BST的中序是升序序列，尝试练习一下非递归中序
func isValidBST(root *TreeNode) bool {
	var last int
	inited := false

	var stack []*TreeNode
	current := root
	for current != nil {
		stack = append(stack, current)
		current = current.Left
	}
	for len(stack) != 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !inited {
			inited = true
		} else {
			if current.Val <= last {
				return false
			}
		}
		last = current.Val

		current = current.Right
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
	}
	return true
}
