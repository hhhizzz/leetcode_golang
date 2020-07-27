package _114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左子树的最右端点接到右子树的最右端点，再返回调节后的节点的最右端点
func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftRight := helper(root.Left)
	rightRight := helper(root.Right)

	if leftRight != nil {
		root.Left, root.Right, leftRight.Right = nil, root.Left, root.Right
	}

	if rightRight != nil {
		return rightRight
	}
	if leftRight != nil {
		return leftRight
	}
	return root
}

func flatten(root *TreeNode) {
	helper(root)
}
