package _98

import "testing"

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{Val: 0}
	//root.Left = &TreeNode{Val: 1}
	//root.Right = &TreeNode{Val: 4}
	println(isValidBST(root))
}
