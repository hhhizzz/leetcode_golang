package _105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	value := preorder[0]
	root := &TreeNode{Val: value}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == value {
			root.Left = buildTree(preorder[1:1+i], inorder[:i])
			root.Right = buildTree(preorder[1+i:], inorder[i+1:])
			break
		}
	}
	return root
}
